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
import { Form, Dropdown, Badge, Alert } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { getHierarchicalTags } from '@/services';
import type { HierarchicalTagItem } from '@/common/interface';

import './index.scss';

interface IProps {
  value?: string;
  onChange?: (tagPath: string, selectedTag: HierarchicalTagItem | null) => void;
  isInvalid?: boolean;
  errMsg?: string;
  required?: boolean;
}

const HierarchicalTagDropdown: React.FC<IProps> = ({
  value = '',
  onChange,
  isInvalid = false,
  errMsg = '',
  required = false,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'hierarchical_tag_dropdown' });
  const [tagLevels, setTagLevels] = useState<HierarchicalTagItem[][]>([]);
  const [selectedPath, setSelectedPath] = useState<HierarchicalTagItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>('');

  // Load root level tags on component mount
  useEffect(() => {
    loadTags('');
  }, []);

  // Parse existing value on mount or when value changes
  useEffect(() => {
    if (value && value !== getCurrentPath()) {
      // TODO: Parse existing path and reconstruct selectedPath
      // This would require additional API to resolve path to tag hierarchy
    }
  }, [value]);

  const loadTags = async (parentId: string, level: number = 0) => {
    setLoading(true);
    setError('');
    try {
      const response = await getHierarchicalTags({ parent_id: parentId });
      const newTagLevels = [...tagLevels];
      newTagLevels[level] = response.tags;
      
      // Clear levels after the current one
      newTagLevels.splice(level + 1);
      
      setTagLevels(newTagLevels);
    } catch (err) {
      console.error('Failed to load hierarchical tags:', err);
      setError('Failed to load tags. Please try again.');
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
    }

    // Always call onChange with current path
    const currentPath = newSelectedPath.map(t => `#${t.display_name}`).join('');
    const finalTag = newSelectedPath[newSelectedPath.length - 1];
    onChange?.(currentPath, finalTag);
  };

  const handleClear = () => {
    setSelectedPath([]);
    setTagLevels(tagLevels.slice(0, 1)); // Keep only root level
    onChange?.('', null);
  };

  const getCurrentPath = () => {
    return selectedPath.map(t => `#${t.display_name}`).join('');
  };

  const canSubmit = selectedPath.length > 0 && !selectedPath[selectedPath.length - 1]?.has_children;

  return (
    <div className="hierarchical-tag-dropdown">
      {error && (
        <Alert variant="danger" className="mb-2">
          {error}
        </Alert>
      )}

      {/* Display current selection */}
      {selectedPath.length > 0 && (
        <div className="current-selection mb-2">
          <div className="d-flex align-items-center">
            <Badge bg="primary" className="me-2">
              {getCurrentPath()}
            </Badge>
            <button
              type="button"
              className="btn btn-sm btn-outline-secondary"
              onClick={handleClear}
              title="Clear selection"
            >
              ✕
            </button>
          </div>
          {!canSubmit && (
            <small className="text-warning">
              {t('select_leaf_node', { defaultValue: 'Please select a specific category (leaf node)' })}
            </small>
          )}
        </div>
      )}

      {/* Cascading dropdowns */}
      <div className="dropdown-levels">
        {tagLevels.map((tags, level) => (
          <div key={level} className="dropdown-level mb-2">
            <Form.Label className="small">
              {level === 0 
                ? t('select_main_category', { defaultValue: 'Main Category' })
                : t('select_subcategory', { defaultValue: `Level ${level + 1}` })
              }
              {level === 0 && required && <span className="text-danger">*</span>}
            </Form.Label>
            <Dropdown>
              <Dropdown.Toggle 
                variant="outline-secondary" 
                className={`w-100 text-start ${isInvalid && level === 0 ? 'is-invalid' : ''}`}
                disabled={loading}
              >
                {selectedPath[level]?.display_name || t('choose_option', { defaultValue: 'Choose...' })}
              </Dropdown.Toggle>
              <Dropdown.Menu className="w-100 dropdown-menu-scrollable">
                {tags.length === 0 ? (
                  <Dropdown.Item disabled>
                    {loading ? 'Loading...' : 'No options available'}
                  </Dropdown.Item>
                ) : (
                  tags.map((tag) => (
                    <Dropdown.Item
                      key={tag.id}
                      onClick={() => handleTagSelect(tag, level)}
                      active={selectedPath[level]?.id === tag.id}
                    >
                      <div className="d-flex justify-content-between align-items-center">
                        <div>
                          <div>{tag.display_name}</div>
                          {tag.description && (
                            <small className="text-muted">{tag.description}</small>
                          )}
                        </div>
                        {tag.has_children && (
                          <i className="bi bi-chevron-right text-muted"></i>
                        )}
                      </div>
                    </Dropdown.Item>
                  ))
                )}
              </Dropdown.Menu>
            </Dropdown>
          </div>
        ))}
      </div>

      {/* Validation message */}
      {isInvalid && errMsg && (
        <div className="invalid-feedback d-block">
          {errMsg}
        </div>
      )}

      {/* Help text */}
      <Form.Text className="text-muted">
        {t('help_text', { 
          defaultValue: 'Select a hierarchical category for your question. Navigate through levels to find the most specific category.' 
        })}
      </Form.Text>
    </div>
  );
};

export default HierarchicalTagDropdown;
