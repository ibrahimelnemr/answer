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

package migrations

import (
	"context"
	"time"

	"xorm.io/xorm"
)

func addHierarchicalTags(ctx context.Context, x *xorm.Engine) error {
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

	type QuestionHierarchicalTagRel struct {
		ID                  int64     `xorm:"not null pk autoincr BIGINT(20) id"`
		CreatedAt           time.Time `xorm:"created TIMESTAMP created_at"`
		QuestionID          string    `xorm:"not null INDEX BIGINT(20) question_id"`
		HierarchicalTagID   string    `xorm:"not null INDEX BIGINT(20) hierarchical_tag_id"`
		HierarchicalTagPath string    `xorm:"not null TEXT hierarchical_tag_path"`
		Status              int       `xorm:"not null default 1 INT(11) status"`
	}

	err := x.Context(ctx).Sync(new(HierarchicalTag), new(QuestionHierarchicalTagRel))
	if err != nil {
		return err
	}

	// Insert sample hierarchical tags
	sampleTags := []HierarchicalTag{
		// Root level tags
		{
			ID:          "1001",
			Name:        "customer",
			SlugName:    "customer",
			DisplayName: "Customer",
			ParentID:    "",
			Level:       0,
			Path:        "#Customer",
			Description: "Customer related topics",
			Status:      1,
			SortOrder:   1,
		},
		{
			ID:          "1002",
			Name:        "internal",
			SlugName:    "internal",
			DisplayName: "Internal",
			ParentID:    "",
			Level:       0,
			Path:        "#Internal",
			Description: "Internal topics",
			Status:      1,
			SortOrder:   2,
		},
		// Customer sub-categories
		{
			ID:          "2001",
			Name:        "fullstack",
			SlugName:    "fullstack",
			DisplayName: "FullStack",
			ParentID:    "1001",
			Level:       1,
			Path:        "#Customer#FullStack",
			Description: "Full stack development",
			Status:      1,
			SortOrder:   1,
		},
		{
			ID:          "2002",
			Name:        "backend",
			SlugName:    "backend",
			DisplayName: "Backend",
			ParentID:    "1001",
			Level:       1,
			Path:        "#Customer#Backend",
			Description: "Backend development",
			Status:      1,
			SortOrder:   2,
		},
		{
			ID:          "2003",
			Name:        "frontend",
			SlugName:    "frontend",
			DisplayName: "Frontend",
			ParentID:    "1001",
			Level:       1,
			Path:        "#Customer#Frontend",
			Description: "Frontend development",
			Status:      1,
			SortOrder:   3,
		},
		{
			ID:          "2004",
			Name:        "mobile",
			SlugName:    "mobile",
			DisplayName: "Mobile",
			ParentID:    "1001",
			Level:       1,
			Path:        "#Customer#Mobile",
			Description: "Mobile development",
			Status:      1,
			SortOrder:   4,
		},
		// Backend sub-categories
		{
			ID:          "3001",
			Name:        "nodejs",
			SlugName:    "nodejs",
			DisplayName: "Node.js",
			ParentID:    "2002",
			Level:       2,
			Path:        "#Customer#Backend#Node.js",
			Description: "Node.js backend development",
			Status:      1,
			SortOrder:   1,
		},
		{
			ID:          "3002",
			Name:        "java",
			SlugName:    "java",
			DisplayName: "Java",
			ParentID:    "2002",
			Level:       2,
			Path:        "#Customer#Backend#Java",
			Description: "Java backend development",
			Status:      1,
			SortOrder:   2,
		},
		{
			ID:          "3003",
			Name:        "python",
			SlugName:    "python",
			DisplayName: "Python",
			ParentID:    "2002",
			Level:       2,
			Path:        "#Customer#Backend#Python",
			Description: "Python backend development",
			Status:      1,
			SortOrder:   3,
		},
		// FullStack sub-categories
		{
			ID:          "3004",
			Name:        "java-fullstack",
			SlugName:    "java-fullstack",
			DisplayName: "Java",
			ParentID:    "2001",
			Level:       2,
			Path:        "#Customer#FullStack#Java",
			Description: "Java full stack development",
			Status:      1,
			SortOrder:   1,
		},
	}

	for _, tag := range sampleTags {
		_, err = x.Context(ctx).Insert(&tag)
		if err != nil {
			return err
		}
	}

	return nil
}
