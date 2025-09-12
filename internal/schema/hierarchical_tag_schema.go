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

package schema

// GetHierarchicalTagsReq request for getting hierarchical tags
type GetHierarchicalTagsReq struct {
	ParentID string `form:"parent_id" json:"parent_id"`
}

// HierarchicalTagItem represents a single hierarchical tag
type HierarchicalTagItem struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	SlugName    string                 `json:"slug_name"`
	DisplayName string                 `json:"display_name"`
	ParentID    string                 `json:"parent_id,omitempty"`
	Level       int                    `json:"level"`
	Path        string                 `json:"path"`
	Description string                 `json:"description,omitempty"`
	HasChildren bool                   `json:"has_children"`
	Children    []*HierarchicalTagItem `json:"children,omitempty"`
}

// GetHierarchicalTagsResp response for getting hierarchical tags
type GetHierarchicalTagsResp struct {
	Tags []*HierarchicalTagItem `json:"tags"`
}

// CreateHierarchicalTagReq request for creating hierarchical tag
type CreateHierarchicalTagReq struct {
	Name        string `json:"name" validate:"required,max=100"`
	SlugName    string `json:"slug_name" validate:"required,max=100"`
	ParentID    string `json:"parent_id,omitempty"`
	DisplayName string `json:"display_name" validate:"required,max=100"`
	Description string `json:"description,omitempty"`
}

// UpdateHierarchicalTagReq request for updating hierarchical tag
type UpdateHierarchicalTagReq struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,max=100"`
	SlugName    string `json:"slug_name" validate:"required,max=100"`
	DisplayName string `json:"display_name" validate:"required,max=100"`
	Description string `json:"description,omitempty"`
}

// HierarchicalTagPathReq request for hierarchical tag path
type HierarchicalTagPathReq struct {
	TagID string `json:"tag_id" validate:"required"`
}

// HierarchicalTagPathResp response for hierarchical tag path
type HierarchicalTagPathResp struct {
	Path        string                 `json:"path"`
	DisplayPath string                 `json:"display_path"`
	Tags        []*HierarchicalTagItem `json:"tags"`
}
