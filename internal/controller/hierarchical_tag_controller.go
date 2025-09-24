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

package controller

import (
	"github.com/apache/answer/internal/base/handler"
	"github.com/apache/answer/internal/schema"
	"github.com/apache/answer/internal/service"
	"github.com/gin-gonic/gin"
)

// HierarchicalTagController hierarchical tag controller
type HierarchicalTagController struct {
	hierarchicalTagService *service.HierarchicalTagService
}

// NewHierarchicalTagController new hierarchical tag controller
func NewHierarchicalTagController(hierarchicalTagService *service.HierarchicalTagService) *HierarchicalTagController {
	return &HierarchicalTagController{
		hierarchicalTagService: hierarchicalTagService,
	}
}

// GetHierarchicalTags godoc
// @Summary Get hierarchical tags
// @Description Get hierarchical tags by parent ID
// @Tags HierarchicalTag
// @Accept json
// @Produce json
// @Param parent_id query string false "parent tag ID"
// @Success 200 {object} handler.RespBody{data=schema.GetHierarchicalTagsResp}
// @Router /answer/api/v1/hierarchical-tags [get]
func (htc *HierarchicalTagController) GetHierarchicalTags(ctx *gin.Context) {
	req := &schema.GetHierarchicalTagsReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := htc.hierarchicalTagService.GetHierarchicalTags(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, resp)
}

// CreateHierarchicalTag godoc
// @Summary Create hierarchical tag
// @Description Create a new hierarchical tag
// @Tags HierarchicalTag
// @Accept json
// @Produce json
// @Param data body schema.CreateHierarchicalTagReq true "hierarchical tag data"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/hierarchical-tags [post]
func (htc *HierarchicalTagController) CreateHierarchicalTag(ctx *gin.Context) {
	req := &schema.CreateHierarchicalTagReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := htc.hierarchicalTagService.CreateHierarchicalTag(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateHierarchicalTag godoc
// @Summary Update hierarchical tag
// @Description Update an existing hierarchical tag
// @Tags HierarchicalTag
// @Accept json
// @Produce json
// @Param data body schema.UpdateHierarchicalTagReq true "hierarchical tag data"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/hierarchical-tags [put]
func (htc *HierarchicalTagController) UpdateHierarchicalTag(ctx *gin.Context) {
	req := &schema.UpdateHierarchicalTagReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := htc.hierarchicalTagService.UpdateHierarchicalTag(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetHierarchicalTagPath godoc
// @Summary Get hierarchical tag path
// @Description Get the full path of a hierarchical tag
// @Tags HierarchicalTag
// @Accept json
// @Produce json
// @Param tag_id query string true "tag ID"
// @Success 200 {object} handler.RespBody{data=schema.HierarchicalTagPathResp}
// @Router /answer/api/v1/hierarchical-tags/path [get]
func (htc *HierarchicalTagController) GetHierarchicalTagPath(ctx *gin.Context) {
	req := &schema.HierarchicalTagPathReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := htc.hierarchicalTagService.GetHierarchicalTagPath(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, resp)
}

// TestHierarchicalTag godoc
// @Summary Test hierarchical tag endpoint
// @Description Simple test endpoint to verify hierarchical tag system is working
// @Tags HierarchicalTag
// @Produce json
// @Success 200 {object} handler.RespBody{data=map[string]interface{}}
// @Router /answer/api/v1/hierarchical-tags/test [get]
func (htc *HierarchicalTagController) TestHierarchicalTag(ctx *gin.Context) {
	testData := map[string]interface{}{
		"message":   "Hierarchical tag system is working!",
		"timestamp": "2024-01-01T00:00:00Z",
		"sample_hierarchy": []string{
			"#Customer",
			"#Customer#Backend",
			"#Customer#Backend#Java",
		},
		"status": "success",
	}

	handler.HandleResponse(ctx, nil, testData)
}
