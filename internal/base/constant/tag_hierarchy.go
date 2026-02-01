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

package constant

import "strings"

const (
	HierarchicalTagDelimiter    = "."
	HierarchicalTagSegmentCount = 3
)

var (
	HierarchicalOfferings = []string{"customer", "internal", "cloud", "ai-data"}
	HierarchicalSpecializations = map[string][]string{
		"customer": {"fullstack-development", "backend-development", "frontend-development", "mobile-app-development", "devops", "qa-testing", "ui-ux-design", "technical-support"},
		"internal": {"fullstack-development", "backend-development", "frontend-development", "data-engineering", "devops", "infrastructure", "security", "project-management"},
		"cloud":   {"cloud-architecture", "cloud-migration", "devops", "infrastructure-as-code", "serverless", "container-orchestration", "cloud-security", "cost-optimization"},
		"ai-data": {"genai-development", "machine-learning", "data-engineering", "data-science", "data-analytics", "mlops", "nlp", "computer-vision"},
	}
	HierarchicalTopics = map[string][]string{
		"fullstack-development": {"java", "python", "javascript", "typescript", "react", "angular", "vue-js", "node-js", "spring-boot", "django", "fastapi", "next-js", "express-js"},
		"backend-development":   {"java", "python", "golang", "node-js", "spring-boot", "micronaut", "django", "fastapi", "quarkus", "ktor"},
		"frontend-development":  {"react", "angular", "vue-js", "next-js", "svelte", "typescript", "storybook", "tailwind", "web-components"},
		"mobile-app-development": {"swift", "kotlin", "react-native", "flutter", "objective-c", "jetpack-compose", "swiftui", "capacitor"},
		"devops":                {"ci-cd", "kubernetes", "docker", "terraform", "ansible", "monitoring", "gitops", "helm"},
		"qa-testing":            {"automation", "selenium", "cypress", "playwright", "junit", "integration-testing", "performance-testing"},
		"ui-ux-design":          {"design-systems", "figma", "adobe-xd", "prototyping", "accessibility", "usability-testing", "motion-design"},
		"technical-support":     {"incident-response", "troubleshooting", "observability", "knowledge-base", "runbooks", "customer-communication"},
		"data-engineering":      {"spark", "kafka", "airflow", "dbt", "etl", "warehousing", "delta-lake", "bigquery"},
		"infrastructure":        {"networking", "virtualization", "observability", "backup", "disaster-recovery", "high-availability", "edge-computing"},
		"security":              {"authentication", "authorization", "oauth", "jwt", "encryption", "penetration-testing", "security-auditing"},
		"project-management":    {"agile", "scrum", "kanban", "roadmapping", "stakeholder-management", "capacity-planning"},
		"cloud-architecture":    {"aws", "azure", "gcp", "terraform", "cloudformation", "architecture-design", "microservices", "serverless", "high-availability", "scalability"},
		"cloud-migration":       {"assessment", "refactor", "lift-and-shift", "hybrid-cloud", "cutover", "data-synchronization"},
		"infrastructure-as-code": {"terraform", "cloudformation", "pulumi", "ansible", "crossplane", "policy-as-code"},
		"serverless":            {"lambda", "functions", "event-driven", "api-gateway", "step-functions", "service-mesh"},
		"container-orchestration": {"kubernetes", "ecs", "aks", "service-mesh", "istio", "helm", "argo"},
		"cloud-security":        {"iam", "guardrails", "compliance", "zero-trust", "network-segmentation", "key-management"},
		"cost-optimization":     {"finops", "autoscaling", "rightsizing", "reserved-instances", "savings-plans", "unit-economics"},
		"genai-development":     {"llms", "prompt-engineering", "rag", "vector-databases", "langchain", "model-evaluation"},
		"machine-learning":      {"python", "tensorflow", "pytorch", "scikit-learn", "keras", "xgboost", "neural-networks", "deep-learning", "model-training", "feature-engineering"},
		"data-science":          {"python", "pandas", "numpy", "statistics", "visualization", "experimentation", "notebooks"},
		"data-analytics":        {"sql", "power-bi", "tableau", "lookml", "dashboards", "business-metrics", "storytelling"},
		"mlops":                 {"model-deployment", "monitoring", "feature-store", "pipelines", "model-registry", "drift-detection"},
		"nlp":                   {"transformers", "tokenization", "embeddings", "sentiment-analysis", "summarization", "speech-to-text"},
		"computer-vision":       {"opencv", "object-detection", "segmentation", "image-processing", "3d-vision", "edge-inference"},
	}
)

var (
	hierarchicalSlugOrder   = make(map[string]int)
	hierarchicalSlugSequence []string
)

func init() {
	index := 0
	for _, offering := range HierarchicalOfferings {
		specializations := HierarchicalSpecializations[offering]
		for _, specialization := range specializations {
			topics := HierarchicalTopics[specialization]
			for _, topic := range topics {
				slug := strings.ToLower(offering + HierarchicalTagDelimiter + specialization + HierarchicalTagDelimiter + topic)
				hierarchicalSlugSequence = append(hierarchicalSlugSequence, slug)
				hierarchicalSlugOrder[slug] = index
				index++
			}
		}
	}
}

// NormalizeHierarchicalSlug lowercases and sanitizes a slug for comparison.
func NormalizeHierarchicalSlug(slug string) string {
	slug = strings.TrimSpace(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ToLower(slug)
	return slug
}

// IsValidHierarchicalSlug verifies if the slug matches the predefined hierarchy.
func IsValidHierarchicalSlug(slug string) bool {
	slug = NormalizeHierarchicalSlug(slug)
	parts := strings.Split(slug, HierarchicalTagDelimiter)
	if len(parts) != HierarchicalTagSegmentCount {
		return false
	}
	offering := parts[0]
	specialization := parts[1]
	topic := parts[2]
	if !containsString(HierarchicalOfferings, offering) {
		return false
	}
	specializationList, ok := HierarchicalSpecializations[offering]
	if !ok || !containsString(specializationList, specialization) {
		return false
	}
	topicList, ok := HierarchicalTopics[specialization]
	if !ok || !containsString(topicList, topic) {
		return false
	}
	return true
}

// HierarchicalSlugOrder returns the deterministic sort order for a slug.
func HierarchicalSlugOrder(slug string) int {
	slug = NormalizeHierarchicalSlug(slug)
	if order, ok := hierarchicalSlugOrder[slug]; ok {
		return order
	}
	return len(hierarchicalSlugOrder)
}

// HierarchicalSlugSequence returns a copy of the ordered slug list used for sorting.
func HierarchicalSlugSequenceCopy() []string {
	result := make([]string, len(hierarchicalSlugSequence))
	copy(result, hierarchicalSlugSequence)
	return result
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
