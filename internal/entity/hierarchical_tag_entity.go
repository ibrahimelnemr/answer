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

package entity

import "time"

const (
	HierarchicalTagStatusAvailable = 1
	HierarchicalTagStatusDeleted   = 10
)

// HierarchicalTag represents a hierarchical tag structure
type HierarchicalTag struct {
	ID          string    `xorm:"not null pk comment('hierarchical_tag_id') BIGINT(20) id"`
	CreatedAt   time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt   time.Time `xorm:"updated TIMESTAMP updated_at"`
	Name        string    `xorm:"not null VARCHAR(100) name"`
	SlugName    string    `xorm:"not null unique VARCHAR(100) slug_name"`
	ParentID    string    `xorm:"default null BIGINT(20) parent_id"`
	Level       int       `xorm:"not null default 0 INT(11) level"`
	Path        string    `xorm:"not null TEXT path"`
	DisplayName string    `xorm:"not null VARCHAR(100) display_name"`
	Description string    `xorm:"TEXT description"`
	Status      int       `xorm:"not null default 1 INT(11) status"`
	SortOrder   int       `xorm:"not null default 0 INT(11) sort_order"`
}

// TableName hierarchical tag table name
func (HierarchicalTag) TableName() string {
	return "hierarchical_tag"
}

// QuestionHierarchicalTagRel represents the relationship between questions and hierarchical tags
type QuestionHierarchicalTagRel struct {
	ID                  int64     `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt           time.Time `xorm:"created TIMESTAMP created_at"`
	QuestionID          string    `xorm:"not null INDEX BIGINT(20) question_id"`
	HierarchicalTagID   string    `xorm:"not null INDEX BIGINT(20) hierarchical_tag_id"`
	HierarchicalTagPath string    `xorm:"not null TEXT hierarchical_tag_path"`
	Status              int       `xorm:"not null default 1 INT(11) status"`
}

// TableName question hierarchical tag relation table name
func (QuestionHierarchicalTagRel) TableName() string {
	return "question_hierarchical_tag_rel"
}

// HierarchicalTagWithChildren represents a hierarchical tag with its children
type HierarchicalTagWithChildren struct {
	HierarchicalTag
	Children []*HierarchicalTagWithChildren `json:"children,omitempty"`
}
