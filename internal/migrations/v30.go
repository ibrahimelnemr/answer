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
	"fmt"
	"strings"

	"github.com/apache/answer/internal/base/data"
	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/repo/unique"
	"xorm.io/xorm"
)

var hierarchicalOfferings = []string{"customer", "internal", "cloud", "ai-data"}

var hierarchicalSpecializations = map[string][]string{
	"customer": {"fullstack-development", "backend-development", "frontend-development", "mobile-app-development", "devops", "qa-testing", "ui-ux-design", "technical-support"},
	"internal": {"fullstack-development", "backend-development", "frontend-development", "data-engineering", "devops", "infrastructure", "security", "project-management"},
	"cloud":    {"cloud-architecture", "cloud-migration", "devops", "infrastructure-as-code", "serverless", "container-orchestration", "cloud-security", "cost-optimization"},
	"ai-data":  {"genai-development", "machine-learning", "data-engineering", "data-science", "data-analytics", "mlops", "nlp", "computer-vision"},
}

var hierarchicalTopics = map[string][]string{
	"fullstack-development":   {"java", "python", "javascript", "typescript", "react", "angular", "vue-js", "node-js", "spring-boot", "django", "fastapi", "next-js", "express-js"},
	"backend-development":     {"java", "python", "golang", "node-js", "spring-boot", "micronaut", "django", "fastapi", "quarkus", "ktor"},
	"frontend-development":    {"react", "angular", "vue-js", "next-js", "svelte", "typescript", "storybook", "tailwind", "web-components"},
	"mobile-app-development":  {"swift", "kotlin", "react-native", "flutter", "objective-c", "jetpack-compose", "swiftui", "capacitor"},
	"devops":                  {"ci-cd", "kubernetes", "docker", "terraform", "ansible", "monitoring", "gitops", "helm"},
	"qa-testing":              {"automation", "selenium", "cypress", "playwright", "junit", "integration-testing", "performance-testing"},
	"ui-ux-design":            {"design-systems", "figma", "adobe-xd", "prototyping", "accessibility", "usability-testing", "motion-design"},
	"technical-support":       {"incident-response", "troubleshooting", "observability", "knowledge-base", "runbooks", "customer-communication"},
	"data-engineering":        {"spark", "kafka", "airflow", "dbt", "etl", "warehousing", "delta-lake", "bigquery"},
	"infrastructure":          {"networking", "virtualization", "observability", "backup", "disaster-recovery", "high-availability", "edge-computing"},
	"security":                {"authentication", "authorization", "oauth", "jwt", "encryption", "penetration-testing", "security-auditing"},
	"project-management":      {"agile", "scrum", "kanban", "roadmapping", "stakeholder-management", "capacity-planning"},
	"cloud-architecture":      {"aws", "azure", "gcp", "terraform", "cloudformation", "architecture-design", "microservices", "serverless", "high-availability", "scalability"},
	"cloud-migration":         {"assessment", "refactor", "lift-and-shift", "hybrid-cloud", "cutover", "data-synchronization"},
	"infrastructure-as-code":  {"terraform", "cloudformation", "pulumi", "ansible", "crossplane", "policy-as-code"},
	"serverless":              {"lambda", "functions", "event-driven", "api-gateway", "step-functions", "service-mesh"},
	"container-orchestration": {"kubernetes", "ecs", "aks", "service-mesh", "istio", "helm", "argo"},
	"cloud-security":          {"iam", "guardrails", "compliance", "zero-trust", "network-segmentation", "key-management"},
	"cost-optimization":       {"finops", "autoscaling", "rightsizing", "reserved-instances", "savings-plans", "unit-economics"},
	"genai-development":       {"llms", "prompt-engineering", "rag", "vector-databases", "langchain", "model-evaluation"},
	"machine-learning":        {"python", "tensorflow", "pytorch", "scikit-learn", "keras", "xgboost", "neural-networks", "deep-learning", "model-training", "feature-engineering"},
	"data-science":            {"python", "pandas", "numpy", "statistics", "visualization", "experimentation", "notebooks"},
	"data-analytics":          {"sql", "power-bi", "tableau", "lookml", "dashboards", "business-metrics", "storytelling"},
	"mlops":                   {"model-deployment", "monitoring", "feature-store", "pipelines", "model-registry", "drift-detection"},
	"nlp":                     {"transformers", "tokenization", "embeddings", "sentiment-analysis", "summarization", "speech-to-text"},
	"computer-vision":         {"opencv", "object-detection", "segmentation", "image-processing", "3d-vision", "edge-inference"},
}

func titleCaseTopic(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

func seedHierarchicalTags(ctx context.Context, x *xorm.Engine) error {
	uniqueIDRepo := unique.NewUniqueIDRepo(&data.Data{DB: x})

	for _, offering := range hierarchicalOfferings {
		specs, ok := hierarchicalSpecializations[offering]
		if !ok {
			continue
		}
		for _, spec := range specs {
			topics, ok := hierarchicalTopics[spec]
			if !ok {
				continue
			}
			for _, topic := range topics {
				slug := fmt.Sprintf("%s.%s.%s", offering, spec, topic)

				// Skip if tag already exists
				exists, err := x.Context(ctx).Where("slug_name = ?", slug).Exist(&entity.Tag{})
				if err != nil {
					return fmt.Errorf("check tag existence for %s: %w", slug, err)
				}
				if exists {
					continue
				}

				tagID, err := uniqueIDRepo.GenUniqueIDStr(ctx, entity.Tag{}.TableName())
				if err != nil {
					return fmt.Errorf("generate unique id for tag %s: %w", slug, err)
				}

				tag := &entity.Tag{
					ID:           tagID,
					SlugName:     slug,
					DisplayName:  titleCaseTopic(topic),
					OriginalText: fmt.Sprintf("Tag for %s in %s/%s.", titleCaseTopic(topic), offering, spec),
					ParsedText:   fmt.Sprintf("<p>Tag for %s in %s/%s.</p>", titleCaseTopic(topic), offering, spec),
					UserID:       "1",
					Status:       entity.TagStatusAvailable,
					RevisionID:   "0",
				}

				if _, err := x.Context(ctx).Insert(tag); err != nil {
					return fmt.Errorf("insert tag %s: %w", slug, err)
				}
			}
		}
	}
	return nil
}
