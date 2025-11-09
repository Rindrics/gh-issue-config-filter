package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
	"time"
)

func outputJSON(ctx context.Context, issuesToCreate IssuesToCreate, defaults Defaults, ghClient GitHubClient) error {
	output := make([]IssueOutput, 0, len(issuesToCreate.Issues))

	for _, issue := range issuesToCreate.Issues {
		// Get target repo
		repo, err := issue.GetTargetRepo(defaults)
		if err != nil {
			return fmt.Errorf("failed to get target repo for issue %s: %w", issue.Name, err)
		}

		// Get project ID
		projectID := defaults.ProjectID
		if issue.ProjectID != nil {
			projectID = *issue.ProjectID
		}

		// Get project fields
		projectFields, err := ghClient.GetProjectFields(ctx, projectID, repo.Owner)
		if err != nil {
			return fmt.Errorf("failed to get project fields for issue %s: %w", issue.Name, err)
		}

		// Create field map for quick lookup
		fieldMap := make(map[string]ProjectField)
		for _, field := range projectFields {
			fieldMap[field.Name] = field
		}

		// Build field_updates array
		fieldUpdates := make([]FieldUpdate, 0, len(issue.Fields))
		for fieldName, fieldValue := range issue.Fields {
			field, exists := fieldMap[fieldName]
			if !exists {
				return fmt.Errorf("field '%s' not found in project for issue %s", fieldName, issue.Name)
			}

			fieldUpdate := FieldUpdate{
				FieldID:   field.ID,
				FieldType: field.DataType,
			}

			// Handle different field types
			switch field.DataType {
			case "TEXT", "NUMBER":
				fieldUpdate.Value = &fieldValue
			case "SINGLE_SELECT":
				// Find option ID by name
				var optionID *string
				for _, opt := range field.Options {
					if opt.Name == fieldValue {
						optionID = &opt.ID
						break
					}
				}
				if optionID == nil {
					return fmt.Errorf("option '%s' not found in field '%s' for issue %s", fieldValue, fieldName, issue.Name)
				}
				fieldUpdate.OptionID = optionID
			default:
				return fmt.Errorf("unsupported field type '%s' for field '%s' in issue %s", field.DataType, fieldName, issue.Name)
			}

			fieldUpdates = append(fieldUpdates, fieldUpdate)
		}

		// Generate title from name and title_suffix
		title, err := expandTitleSuffix(issue.Name, issue.TitleSuffix)
		if err != nil {
			return fmt.Errorf("failed to expand title for issue %s: %w", issue.Name, err)
		}

		item := IssueOutput{
			Name:         issue.Name,
			Title:        title,
			TemplateFile: issue.TemplateFile,
			ProjectID:    issue.ProjectID,
			TargetRepo:   issue.TargetRepo,
			FieldUpdates: fieldUpdates,
		}

		output = append(output, item)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}

// expandTitleSuffix expands template variables in title_suffix and returns the final title.
// If titleSuffix is nil or empty, returns the name as-is.
// Supported template functions:
//   - {{Date}} - Current date in YYYY-MM-DD format
//   - {{Year}} - Current year (e.g., 2025)
//   - {{Month}} - Current month (e.g., 01)
//   - {{YearMonth}} - Current year and month in YYYY-MM format
func expandTitleSuffix(name string, titleSuffix *string) (string, error) {
	if titleSuffix == nil || *titleSuffix == "" {
		return name, nil
	}

	now := time.Now()
	funcMap := template.FuncMap{
		"Date": func() string {
			return now.Format("2006-01-02")
		},
		"Year": func() string {
			return now.Format("2006")
		},
		"Month": func() string {
			return now.Format("01")
		},
		"YearMonth": func() string {
			return now.Format("2006-01")
		},
	}

	tmpl, err := template.New("title").Funcs(funcMap).Parse(*titleSuffix)
	if err != nil {
		return "", fmt.Errorf("failed to parse title_suffix template: %w", err)
	}

	var buf bytes.Buffer
	// Execute with empty data since we're using functions, not data fields
	if err := tmpl.Execute(&buf, struct{}{}); err != nil {
		return "", fmt.Errorf("failed to execute title_suffix template: %w", err)
	}

	expandedSuffix := buf.String()
	if expandedSuffix == "" {
		return name, nil
	}

	return fmt.Sprintf("%s %s", name, expandedSuffix), nil
}
