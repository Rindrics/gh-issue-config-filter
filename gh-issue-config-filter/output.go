package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
		
		item := IssueOutput{
			Name:         issue.Name,
			TemplateFile: issue.TemplateFile,
			ProjectID:    issue.ProjectID,
			TargetRepo:   issue.TargetRepo,
			FieldUpdates: fieldUpdates,
		}
		if issue.TitleSuffix != nil {
			item.TitleSuffix = issue.TitleSuffix
		}
		
		output = append(output, item)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}

