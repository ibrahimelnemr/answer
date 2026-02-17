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

import { useMemo } from 'react';
import { Accordion, Badge, Col, Row, Stack } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { useTagHierarchy } from '@/services';
import type * as Type from '@/common/interface';

const TagsHierarchy = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'tags' });
  const { data, isLoading } = useTagHierarchy();

  const offerings = useMemo<Type.TagHierarchyOffering[]>(
    () => data?.offerings || [],
    [data],
  );

  usePageTags({
    title: t('tags', { keyPrefix: 'page_title' }),
  });

  return (
    <Row className="py-4 mb-4">
      <Col xxl={12}>
        <Stack direction="horizontal" className="mb-4" gap={3}>
          <div>
            <h3 className="mb-1">{t('title')}</h3>
            <div className="text-secondary small">Hierarchy view</div>
          </div>
          <div className="ms-auto">
            <Link to="/tags" className="btn btn-outline-secondary btn-sm">
              Back to tags
            </Link>
          </div>
        </Stack>
      </Col>

      <Col xxl={12}>
        {isLoading ? (
          <div className="text-secondary">Loading tag hierarchy...</div>
        ) : offerings.length === 0 ? (
          <div className="text-secondary">No hierarchical tags found.</div>
        ) : (
          <Accordion alwaysOpen>
            {offerings.map((offering, offeringIndex) => (
              <Accordion.Item
                key={offering.name}
                eventKey={String(offeringIndex)}>
                <Accordion.Header>{offering.name}</Accordion.Header>
                <Accordion.Body>
                  {offering.specializations.map((spec) => (
                    <div key={spec.name} className="mb-3">
                      <div className="fw-semibold mb-2">{spec.name}</div>
                      <div className="d-flex flex-wrap gap-2">
                        {spec.topics.map((topic) => {
                          const slug = `${offering.name}.${spec.name}.${topic}`;
                          return (
                            <Badge
                              key={slug}
                              bg="light"
                              text="dark"
                              className="border">
                              <Link
                                className="text-decoration-none text-reset"
                                to={`/tags/${encodeURIComponent(slug)}`}>
                                {topic}
                              </Link>
                            </Badge>
                          );
                        })}
                      </div>
                    </div>
                  ))}
                </Accordion.Body>
              </Accordion.Item>
            ))}
          </Accordion>
        )}
      </Col>
    </Row>
  );
};

export default TagsHierarchy;
