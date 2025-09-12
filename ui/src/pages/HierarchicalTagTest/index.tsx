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

import React, { useState } from 'react';
import { Container, Row, Col, Card, Alert } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { HierarchicalTagDropdown, HierarchicalTagSelector } from '@/components';
import type { HierarchicalTagItem } from '@/common/interface';

const HierarchicalTagTest: React.FC = () => {
  const { t } = useTranslation();
  const [dropdownValue, setDropdownValue] = useState('');
  const [selectedTag, setSelectedTag] = useState<HierarchicalTagItem | null>(null);
  const [showSelector, setShowSelector] = useState(false);
  const [selectorResult, setSelectorResult] = useState('');

  const handleDropdownChange = (tagPath: string, tag: HierarchicalTagItem | null) => {
    setDropdownValue(tagPath);
    setSelectedTag(tag);
  };

  const handleSelectorConfirm = (tag: HierarchicalTagItem, path: string) => {
    setSelectorResult(path);
    setShowSelector(false);
  };

  return (
    <Container className="py-4">
      <Row>
        <Col md={8} className="mx-auto">
          <h1>Hierarchical Tag Components Test</h1>
          
          <Alert variant="info">
            <strong>Testing Instructions:</strong>
            <ul>
              <li>Test the dropdown component below - it should load hierarchical tags dynamically</li>
              <li>Test the modal selector by clicking the button</li>
              <li>Check browser console for any errors</li>
              <li>Verify API calls are being made to <code>/answer/api/v1/hierarchical-tags</code></li>
            </ul>
          </Alert>

          <Card className="mb-4">
            <Card.Header>
              <h5>Hierarchical Tag Dropdown Component</h5>
            </Card.Header>
            <Card.Body>
              <HierarchicalTagDropdown
                value={dropdownValue}
                onChange={handleDropdownChange}
                required
              />
              
              {dropdownValue && (
                <div className="mt-3">
                  <Alert variant="success">
                    <strong>Selected Path:</strong> {dropdownValue}
                    <br />
                    <strong>Selected Tag:</strong> {selectedTag?.display_name} (ID: {selectedTag?.id})
                  </Alert>
                </div>
              )}
            </Card.Body>
          </Card>

          <Card className="mb-4">
            <Card.Header>
              <h5>Hierarchical Tag Selector Modal</h5>
            </Card.Header>
            <Card.Body>
              <button
                className="btn btn-primary"
                onClick={() => setShowSelector(true)}
              >
                Open Tag Selector Modal
              </button>
              
              {selectorResult && (
                <div className="mt-3">
                  <Alert variant="success">
                    <strong>Modal Result:</strong> {selectorResult}
                  </Alert>
                </div>
              )}
            </Card.Body>
          </Card>

          <Card>
            <Card.Header>
              <h5>API Testing</h5>
            </Card.Header>
            <Card.Body>
              <p>
                Open browser developer tools and check the Network tab. 
                You should see API calls to:
              </p>
              <ul>
                <li><code>GET /answer/api/v1/hierarchical-tags</code> - Load root tags</li>
                <li><code>GET /answer/api/v1/hierarchical-tags?parent_id=X</code> - Load child tags</li>
              </ul>
              <p>
                If you see 404 errors, the backend migration may not have run yet.
              </p>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      <HierarchicalTagSelector
        visible={showSelector}
        onClose={() => setShowSelector(false)}
        onConfirm={handleSelectorConfirm}
      />
    </Container>
  );
};

export default HierarchicalTagTest;
