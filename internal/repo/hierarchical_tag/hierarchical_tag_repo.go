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

package hierarchical_tag

import (
	"context"

	"github.com/apache/answer/internal/base/data"
	"github.com/apache/answer/internal/entity"
	"github.com/segmentfault/pacman/errors"
)

// HierarchicalTagRepo hierarchical tag repository
type HierarchicalTagRepo struct {
	data *data.Data
}

// NewHierarchicalTagRepo new repository
func NewHierarchicalTagRepo(data *data.Data) *HierarchicalTagRepo {
	return &HierarchicalTagRepo{
		data: data,
	}
}

// GetByParentID gets hierarchical tags by parent ID
func (hr *HierarchicalTagRepo) GetByParentID(ctx context.Context, parentID string) (tags []*entity.HierarchicalTag, err error) {
	session := hr.data.DB.Context(ctx).Where("status = ?", entity.HierarchicalTagStatusAvailable)
	if parentID == "" || parentID == "0" {
		session = session.Where("parent_id IS NULL OR parent_id = ''")
	} else {
		session = session.Where("parent_id = ?", parentID)
	}

	err = session.OrderBy("sort_order ASC, display_name ASC").Find(&tags)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// GetByID gets hierarchical tag by ID
func (hr *HierarchicalTagRepo) GetByID(ctx context.Context, id string) (tag *entity.HierarchicalTag, exist bool, err error) {
	tag = &entity.HierarchicalTag{}
	exist, err = hr.data.DB.Context(ctx).Where("id = ? AND status = ?", id, entity.HierarchicalTagStatusAvailable).Get(tag)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// GetBySlugName gets hierarchical tag by slug name
func (hr *HierarchicalTagRepo) GetBySlugName(ctx context.Context, slugName string) (tag *entity.HierarchicalTag, exist bool, err error) {
	tag = &entity.HierarchicalTag{}
	exist, err = hr.data.DB.Context(ctx).Where("slug_name = ? AND status = ?", slugName, entity.HierarchicalTagStatusAvailable).Get(tag)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// Create creates a new hierarchical tag
func (hr *HierarchicalTagRepo) Create(ctx context.Context, tag *entity.HierarchicalTag) (err error) {
	// Calculate level and path
	if tag.ParentID != "" && tag.ParentID != "0" {
		parent, exist, err := hr.GetByID(ctx, tag.ParentID)
		if err != nil {
			return err
		}
		if !exist {
			return errors.BadRequest("parent tag not found")
		}
		tag.Level = parent.Level + 1
		tag.Path = parent.Path + "#" + tag.DisplayName
	} else {
		tag.Level = 0
		tag.Path = "#" + tag.DisplayName
	}

	_, err = hr.data.DB.Context(ctx).Insert(tag)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// Update updates hierarchical tag
func (hr *HierarchicalTagRepo) Update(ctx context.Context, tag *entity.HierarchicalTag) (err error) {
	_, err = hr.data.DB.Context(ctx).Where("id = ?", tag.ID).Update(tag)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// GetPath gets the full path of a hierarchical tag
func (hr *HierarchicalTagRepo) GetPath(ctx context.Context, tagID string) (path string, tags []*entity.HierarchicalTag, err error) {
	tag, exist, err := hr.GetByID(ctx, tagID)
	if err != nil {
		return
	}
	if !exist {
		err = errors.BadRequest("tag not found")
		return
	}

	path = tag.Path
	tags = []*entity.HierarchicalTag{tag}

	// Get all parent tags
	currentTag := tag
	for currentTag.ParentID != "" && currentTag.ParentID != "0" {
		parent, exist, err := hr.GetByID(ctx, currentTag.ParentID)
		if err != nil {
			return path, tags, err
		}
		if !exist {
			break
		}
		tags = append([]*entity.HierarchicalTag{parent}, tags...)
		currentTag = parent
	}

	return
}

// HasChildren checks if a hierarchical tag has children
func (hr *HierarchicalTagRepo) HasChildren(ctx context.Context, tagID string) (bool, error) {
	count, err := hr.data.DB.Context(ctx).Where("parent_id = ? AND status = ?", tagID, entity.HierarchicalTagStatusAvailable).Count(&entity.HierarchicalTag{})
	if err != nil {
		return false, errors.InternalServer(err.Error())
	}
	return count > 0, nil
}

// CreateQuestionTagRel creates relationship between question and hierarchical tag
func (hr *HierarchicalTagRepo) CreateQuestionTagRel(ctx context.Context, questionID, tagID string) (err error) {
	// Get the full path for the tag
	path, _, err := hr.GetPath(ctx, tagID)
	if err != nil {
		return err
	}

	rel := &entity.QuestionHierarchicalTagRel{
		QuestionID:          questionID,
		HierarchicalTagID:   tagID,
		HierarchicalTagPath: path,
		Status:              1,
	}

	_, err = hr.data.DB.Context(ctx).Insert(rel)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// GetQuestionTagRels gets hierarchical tag relationships for a question
func (hr *HierarchicalTagRepo) GetQuestionTagRels(ctx context.Context, questionID string) (rels []*entity.QuestionHierarchicalTagRel, err error) {
	err = hr.data.DB.Context(ctx).Where("question_id = ? AND status = ?", questionID, 1).Find(&rels)
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}

// RemoveQuestionTagRels removes all hierarchical tag relationships for a question
func (hr *HierarchicalTagRepo) RemoveQuestionTagRels(ctx context.Context, questionID string) (err error) {
	_, err = hr.data.DB.Context(ctx).Where("question_id = ?", questionID).Update(&entity.QuestionHierarchicalTagRel{Status: 0})
	if err != nil {
		err = errors.InternalServer(err.Error())
	}
	return
}
