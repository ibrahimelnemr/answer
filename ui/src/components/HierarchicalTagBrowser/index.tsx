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

import { FC, useMemo, useState } from 'react';
import { Form, Spinner } from 'react-bootstrap';

import { useTagHierarchy } from '@/services';
import type * as Type from '@/common/interface';

interface IProps {
  onSelect?: (slug: string) => void;
}

const HierarchicalTagBrowser: FC<IProps> = ({ onSelect }) => {
  const { data, isLoading } = useTagHierarchy();
  const [selectedOffering, setSelectedOffering] = useState('');
  const [selectedSpec, setSelectedSpec] = useState('');
  const [selectedTopic, setSelectedTopic] = useState('');

  const offerings = useMemo<Type.TagHierarchyOffering[]>(
    () => data?.offerings || [],
    [data],
  );

  const selectedOfferingNode = useMemo(
    () => offerings.find((o) => o.name === selectedOffering),
    [offerings, selectedOffering],
  );

  const specializations = useMemo<Type.TagHierarchySpecialization[]>(
    () => selectedOfferingNode?.specializations || [],
    [selectedOfferingNode],
  );

  const selectedSpecNode = useMemo(
    () => specializations.find((s) => s.name === selectedSpec),
    [specializations, selectedSpec],
  );

  const topics = useMemo<string[]>(
    () => selectedSpecNode?.topics || [],
    [selectedSpecNode],
  );

  const handleOfferingChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const val = e.target.value;
    setSelectedOffering(val);
    setSelectedSpec('');
    setSelectedTopic('');
  };

  const handleSpecChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const val = e.target.value;
    setSelectedSpec(val);
    setSelectedTopic('');
  };

  const handleTopicChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const val = e.target.value;
    setSelectedTopic(val);
    if (val && selectedOffering && selectedSpec && onSelect) {
      onSelect(`${selectedOffering}.${selectedSpec}.${val}`);
    }
  };

  if (isLoading) {
    return (
      <div className="d-flex align-items-center text-secondary small mt-2 mb-3">
        <Spinner animation="border" size="sm" className="me-2" />
        Loading existing tags...
      </div>
    );
  }

  if (!offerings.length) {
    return (
      <div className="text-secondary small mt-2 mb-3">
        No hierarchical tags found. Enter a slug manually above.
      </div>
    );
  }

  return (
    <div className="mt-2 mb-3 p-3 border rounded bg-light">
      <Form.Label className="small fw-bold mb-2">
        Browse existing tags
      </Form.Label>
      <div className="d-flex gap-2 flex-wrap">
        <Form.Select
          size="sm"
          style={{ maxWidth: '200px' }}
          value={selectedOffering}
          onChange={handleOfferingChange}>
          <option value="">-- Offering --</option>
          {offerings.map((o) => (
            <option key={o.name} value={o.name}>
              {o.name}
            </option>
          ))}
        </Form.Select>

        <Form.Select
          size="sm"
          style={{ maxWidth: '240px' }}
          value={selectedSpec}
          onChange={handleSpecChange}
          disabled={!selectedOffering}>
          <option value="">-- Specialization --</option>
          {specializations.map((s) => (
            <option key={s.name} value={s.name}>
              {s.name}
            </option>
          ))}
        </Form.Select>

        <Form.Select
          size="sm"
          style={{ maxWidth: '200px' }}
          value={selectedTopic}
          onChange={handleTopicChange}
          disabled={!selectedSpec}>
          <option value="">-- Topic --</option>
          {topics.map((t) => (
            <option key={t} value={t}>
              {t}
            </option>
          ))}
        </Form.Select>
      </div>
      {selectedOffering && selectedSpec && selectedTopic && (
        <div className="mt-2 small text-muted">
          Selected:{' '}
          <code>
            {selectedOffering}.{selectedSpec}.{selectedTopic}
          </code>
        </div>
      )}
    </div>
  );
};

export default HierarchicalTagBrowser;
