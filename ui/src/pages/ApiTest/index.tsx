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
import { Container, Card, Button, Alert } from 'react-bootstrap';

const ApiTest: React.FC = () => {
  const [testResult, setTestResult] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>('');

  const testApi = async () => {
    setLoading(true);
    setError('');
    try {
      const response = await fetch('/answer/api/v1/hierarchical-tags/test');
      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }
      const data = await response.json();
      setTestResult(data);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const testHierarchicalTags = async () => {
    setLoading(true);
    setError('');
    try {
      const response = await fetch('/answer/api/v1/hierarchical-tags');
      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }
      const data = await response.json();
      setTestResult(data);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  // Auto-test on component mount
  useEffect(() => {
    testApi();
  }, []);

  return (
    <Container className="py-4">
      <h1>API Test Page</h1>
      <p>This page tests if the hierarchical tag API endpoints are working.</p>

      <Card className="mb-4">
        <Card.Header>
          <h5>Test Results</h5>
        </Card.Header>
        <Card.Body>
          <div className="mb-3">
            <Button 
              onClick={testApi} 
              disabled={loading}
              className="me-2"
              variant="primary"
            >
              {loading ? 'Testing...' : 'Test Simple Endpoint'}
            </Button>
            <Button 
              onClick={testHierarchicalTags} 
              disabled={loading}
              variant="secondary"
            >
              {loading ? 'Testing...' : 'Test Hierarchical Tags'}
            </Button>
          </div>

          {error && (
            <Alert variant="danger">
              <strong>Error:</strong> {error}
              <hr />
              <small>
                If you see a 404 error, the backend routes are not registered properly.<br />
                If you see a network error, make sure the backend is running.
              </small>
            </Alert>
          )}

          {testResult && (
            <Alert variant="success">
              <strong>Success!</strong> API is working.
              <pre style={{ marginTop: '10px', fontSize: '12px' }}>
                {JSON.stringify(testResult, null, 2)}
              </pre>
            </Alert>
          )}
        </Card.Body>
      </Card>

      <Card>
        <Card.Header>
          <h5>Manual Testing</h5>
        </Card.Header>
        <Card.Body>
          <p>You can also test these endpoints manually:</p>
          <ul>
            <li>
              <strong>Test Endpoint:</strong>{' '}
              <a href="/answer/api/v1/hierarchical-tags/test" target="_blank">
                /answer/api/v1/hierarchical-tags/test
              </a>
            </li>
            <li>
              <strong>Hierarchical Tags:</strong>{' '}
              <a href="/answer/api/v1/hierarchical-tags" target="_blank">
                /answer/api/v1/hierarchical-tags
              </a>
            </li>
          </ul>
        </Card.Body>
      </Card>
    </Container>
  );
};

export default ApiTest;

