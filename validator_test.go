package main

import (
	"testing"
)

func TestValidateIssue(t *testing.T) {
	cases := []struct {
		name        string
		issue       IssueToCreate
		expectError bool
	}{
		{
			name: "valid issue",
			issue: IssueToCreate{
				Issue: Issue{
					Name: "test",
				},
			},
			expectError: false,
		},
		{
			name: "invalid - empty name",
			issue: IssueToCreate{
				Issue: Issue{
					Name: "",
				},
			},
			expectError: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIssue(tt.issue)
			if tt.expectError && err == nil {
				t.Errorf("expected %v, got %v", tt.expectError, err)
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
