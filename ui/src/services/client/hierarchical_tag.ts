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

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const getHierarchicalTags = (params: Type.GetHierarchicalTagsReq): Promise<Type.GetHierarchicalTagsResp> => {
  const queryParams = new URLSearchParams();
  if (params.parent_id) {
    queryParams.append('parent_id', params.parent_id);
  }
  
  const apiUrl = `/answer/api/v1/hierarchical-tags?${queryParams.toString()}`;
  return request.get<Type.GetHierarchicalTagsResp>(apiUrl);
};

export const getHierarchicalTagPath = (params: Type.HierarchicalTagPathReq): Promise<Type.HierarchicalTagPathResp> => {
  const apiUrl = `/answer/api/v1/hierarchical-tags/path?tag_id=${params.tag_id}`;
  return request.get<Type.HierarchicalTagPathResp>(apiUrl);
};

export const createHierarchicalTag = (params: {
  name: string;
  slug_name: string;
  parent_id?: string;
  display_name: string;
  description?: string;
}): Promise<void> => {
  return request.post('/answer/api/v1/hierarchical-tags', params);
};

export const updateHierarchicalTag = (params: {
  id: string;
  name: string;
  slug_name: string;
  display_name: string;
  description?: string;
}): Promise<void> => {
  return request.put('/answer/api/v1/hierarchical-tags', params);
};
