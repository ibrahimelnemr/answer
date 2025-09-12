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

package service

import (
	"context"

	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/repo/hierarchical_tag"
	"github.com/apache/answer/internal/schema"
	"github.com/apache/answer/pkg/uid"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// HierarchicalTagService hierarchical tag service
type HierarchicalTagService struct {
	hierarchicalTagRepo *hierarchical_tag.HierarchicalTagRepo
}

// NewHierarchicalTagService new hierarchical tag service
func NewHierarchicalTagService(hierarchicalTagRepo *hierarchical_tag.HierarchicalTagRepo) *HierarchicalTagService {
	return &HierarchicalTagService{
		hierarchicalTagRepo: hierarchicalTagRepo,
	}
}

// GetHierarchicalTags gets hierarchical tags by parent ID
func (hs *HierarchicalTagService) GetHierarchicalTags(ctx context.Context, req *schema.GetHierarchicalTagsReq) (*schema.GetHierarchicalTagsResp, error) {
	tags, err := hs.hierarchicalTagRepo.GetByParentID(ctx, req.ParentID)
	if err != nil {
		return nil, err
	}

	resp := &schema.GetHierarchicalTagsResp{
		Tags: make([]*schema.HierarchicalTagItem, 0, len(tags)),
	}

	for _, tag := range tags {
		hasChildren, err := hs.hierarchicalTagRepo.HasChildren(ctx, tag.ID)
		if err != nil {
			log.Error(err)
			hasChildren = false
		}

		tagItem := &schema.HierarchicalTagItem{
			ID:          tag.ID,
			Name:        tag.Name,
			SlugName:    tag.SlugName,
			DisplayName: tag.DisplayName,
			ParentID:    tag.ParentID,
			Level:       tag.Level,
			Path:        tag.Path,
			Description: tag.Description,
			HasChildren: hasChildren,
		}
		resp.Tags = append(resp.Tags, tagItem)
	}

	return resp, nil
}

// CreateHierarchicalTag creates a new hierarchical tag
func (hs *HierarchicalTagService) CreateHierarchicalTag(ctx context.Context, req *schema.CreateHierarchicalTagReq) error {
	tag := &entity.HierarchicalTag{
		ID:          uid.IDStr(),
		Name:        req.Name,
		SlugName:    req.SlugName,
		ParentID:    req.ParentID,
		DisplayName: req.DisplayName,
		Description: req.Description,
		Status:      entity.HierarchicalTagStatusAvailable,
		SortOrder:   0,
	}

	return hs.hierarchicalTagRepo.Create(ctx, tag)
}

// UpdateHierarchicalTag updates hierarchical tag
func (hs *HierarchicalTagService) UpdateHierarchicalTag(ctx context.Context, req *schema.UpdateHierarchicalTagReq) error {
	tag, exist, err := hs.hierarchicalTagRepo.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest("tag not found")
	}

	tag.Name = req.Name
	tag.SlugName = req.SlugName
	tag.DisplayName = req.DisplayName
	tag.Description = req.Description

	return hs.hierarchicalTagRepo.Update(ctx, tag)
}

// GetHierarchicalTagPath gets the full path of a hierarchical tag
func (hs *HierarchicalTagService) GetHierarchicalTagPath(ctx context.Context, req *schema.HierarchicalTagPathReq) (*schema.HierarchicalTagPathResp, error) {
	path, tags, err := hs.hierarchicalTagRepo.GetPath(ctx, req.TagID)
	if err != nil {
		return nil, err
	}

	resp := &schema.HierarchicalTagPathResp{
		Path:        path,
		DisplayPath: path,
		Tags:        make([]*schema.HierarchicalTagItem, 0, len(tags)),
	}

	for _, tag := range tags {
		hasChildren, err := hs.hierarchicalTagRepo.HasChildren(ctx, tag.ID)
		if err != nil {
			log.Error(err)
			hasChildren = false
		}

		tagItem := &schema.HierarchicalTagItem{
			ID:          tag.ID,
			Name:        tag.Name,
			SlugName:    tag.SlugName,
			DisplayName: tag.DisplayName,
			ParentID:    tag.ParentID,
			Level:       tag.Level,
			Path:        tag.Path,
			Description: tag.Description,
			HasChildren: hasChildren,
		}
		resp.Tags = append(resp.Tags, tagItem)
	}

	return resp, nil
}

// AddHierarchicalTagsToQuestion adds hierarchical tags to a question
func (hs *HierarchicalTagService) AddHierarchicalTagsToQuestion(ctx context.Context, questionID string, tagIDs []string) error {
	// Remove existing relationships
	err := hs.hierarchicalTagRepo.RemoveQuestionTagRels(ctx, questionID)
	if err != nil {
		return err
	}

	// Add new relationships
	for _, tagID := range tagIDs {
		err = hs.hierarchicalTagRepo.CreateQuestionTagRel(ctx, questionID, tagID)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetQuestionHierarchicalTags gets hierarchical tags for a question
func (hs *HierarchicalTagService) GetQuestionHierarchicalTags(ctx context.Context, questionID string) ([]*schema.HierarchicalTagItem, error) {
	rels, err := hs.hierarchicalTagRepo.GetQuestionTagRels(ctx, questionID)
	if err != nil {
		return nil, err
	}

	tags := make([]*schema.HierarchicalTagItem, 0, len(rels))
	for _, rel := range rels {
		tag, exist, err := hs.hierarchicalTagRepo.GetByID(ctx, rel.HierarchicalTagID)
		if err != nil {
			return nil, err
		}
		if !exist {
			continue
		}

		hasChildren, err := hs.hierarchicalTagRepo.HasChildren(ctx, tag.ID)
		if err != nil {
			log.Error(err)
			hasChildren = false
		}

		tagItem := &schema.HierarchicalTagItem{
			ID:          tag.ID,
			Name:        tag.Name,
			SlugName:    tag.SlugName,
			DisplayName: tag.DisplayName,
			ParentID:    tag.ParentID,
			Level:       tag.Level,
			Path:        rel.HierarchicalTagPath,
			Description: tag.Description,
			HasChildren: hasChildren,
		}
		tags = append(tags, tagItem)
	}

	return tags, nil
}
