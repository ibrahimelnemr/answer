/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { getHierarchicalTags } from '@/services';
import type { HierarchicalTagItem } from '@/common/interface';

import './index.scss';

interface IProps {
  visible: boolean;
  onClose: () => void;
  onConfirm: (selectedTag: HierarchicalTagItem, path: string) => void;
}

const HierarchicalTagSelector: React.FC<IProps> = ({
  visible,
  onClose,
  onConfirm,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'hierarchical_tag_selector' });
  const [tagLevels, setTagLevels] = useState<HierarchicalTagItem[][]>([]);
  const [selectedPath, setSelectedPath] = useState<HierarchicalTagItem[]>([]);
  const [loading, setLoading] = useState(false);

  // Load root level tags when modal opens
  useEffect(() => {
    if (visible) {
      loadTags('');
    }
  }, [visible]);

  const loadTags = async (parentId: string, level: number = 0) => {
    setLoading(true);
    try {
      const response = await getHierarchicalTags({ parent_id: parentId });
      const newTagLevels = [...tagLevels];
      newTagLevels[level] = response.tags;
      
      // Clear levels after the current one
      newTagLevels.splice(level + 1);
      
      setTagLevels(newTagLevels);
    } catch (error) {
      console.error('Failed to load hierarchical tags:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleTagSelect = async (tag: HierarchicalTagItem, level: number) => {
    const newSelectedPath = [...selectedPath];
    newSelectedPath[level] = tag;
    
    // Clear path after current level
    newSelectedPath.splice(level + 1);
    
    setSelectedPath(newSelectedPath);

    if (tag.has_children) {
      // Load children for next level
      await loadTags(tag.id, level + 1);
    } else {
      // This is a leaf node, we can confirm selection
      const fullPath = newSelectedPath.map(t => `#${t.display_name}`).join('');
      onConfirm(tag, fullPath);
    }
  };

  const handleConfirm = () => {
    if (selectedPath.length > 0) {
      const lastSelected = selectedPath[selectedPath.length - 1];
      const fullPath = selectedPath.map(t => `#${t.display_name}`).join('');
      onConfirm(lastSelected, fullPath);
    }
  };

  const handleClose = () => {
    setTagLevels([]);
    setSelectedPath([]);
    onClose();
  };

  const canConfirm = selectedPath.length > 0 && !selectedPath[selectedPath.length - 1]?.has_children;

  return (
    <Modal show={visible} onHide={handleClose} size="lg">
      <Modal.Header closeButton>
        <Modal.Title>{t('title', { defaultValue: 'Select Tag' })}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className="hierarchical-tag-selector">
          {/* Breadcrumb showing selected path */}
          {selectedPath.length > 0 && (
            <div className="selected-path mb-3">
              <strong>{t('selected_path', { defaultValue: 'Selected Path' })}: </strong>
              <span className="path-display">
                {selectedPath.map(tag => `#${tag.display_name}`).join('')}
              </span>
            </div>
          )}

          {/* Dropdown levels */}
          <div className="tag-levels">
            {tagLevels.map((tags, level) => (
              <div key={level} className="tag-level mb-3">
                <Form.Label>
                  {level === 0 
                    ? t('select_category', { defaultValue: 'Select Category' })
                    : t('select_subcategory', { defaultValue: `Select Level ${level + 1}` })
                  }
                </Form.Label>
                <Dropdown>
                  <Dropdown.Toggle variant="outline-secondary" className="w-100 text-start">
                    {selectedPath[level]?.display_name || t('choose_option', { defaultValue: 'Choose...' })}
                  </Dropdown.Toggle>
                  <Dropdown.Menu className="w-100">
                    {tags.map((tag) => (
                      <Dropdown.Item
                        key={tag.id}
                        onClick={() => handleTagSelect(tag, level)}
                        active={selectedPath[level]?.id === tag.id}
                      >
                        <div className="d-flex justify-content-between align-items-center">
                          <span>{tag.display_name}</span>
                          {tag.has_children && (
                            <i className="bi bi-chevron-right text-muted"></i>
                          )}
                        </div>
                        {tag.description && (
                          <small className="text-muted d-block">{tag.description}</small>
                        )}
                      </Dropdown.Item>
                    ))}
                  </Dropdown.Menu>
                </Dropdown>
              </div>
            ))}
          </div>

          {loading && (
            <div className="text-center">
              <div className="spinner-border" role="status">
                <span className="visually-hidden">Loading...</span>
              </div>
            </div>
          )}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          {t('cancel', { defaultValue: 'Cancel' })}
        </Button>
        <Button 
          variant="primary" 
          onClick={handleConfirm}
          disabled={!canConfirm}
        >
          {t('confirm', { defaultValue: 'Confirm Selection' })}
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default HierarchicalTagSelector;
