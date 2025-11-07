package main

import (
	"testing"
)

func TestGetIssuesToCreate(t *testing.T) {
	defaults := Defaults{
		ProjectID:  "default_project_id",
		TargetRepo: "default/repo",
	}
	issue1 := Issue{
		Name:           "Issue 1",
		CreationMonths: []Month{1},
	}
	issue2 := Issue{
		Name:           "Issue 2",
		CreationMonths: []Month{2},
	}
	issue1_3 := Issue{
		Name:           "Issue 1_3",
		CreationMonths: []Month{1, 3},
	}
	issue2_4 := Issue{
		Name:           "Issue 2_4",
		CreationMonths: []Month{2, 4},
	}
	issue_project_repo := Issue{
		Name:           "Issue project_repo",
		CreationMonths: []Month{1},
		ProjectID:      "other_project_id",
		TargetRepo:     "other/repo",
	}

	cases := []struct {
		name           string
		config         Config
		month          Month
		issuesToCreate IssuesToCreate
	}{
		{
			name: "No issues",
			config: Config{
				Defaults: defaults,
				Issues:   []Issue{},
			},
			month:          Month(1),
			issuesToCreate: IssuesToCreate{},
		},
		{
			name: "One issue",
			config: Config{
				Defaults: defaults,
				Issues:   []Issue{issue1},
			},
			month: Month(1),
			issuesToCreate: IssuesToCreate{
				Issues: []IssueToCreate{
					{
						Issue:      issue1,
						ProjectID:  "default_project_id",
						TargetRepo: "default/repo",
					},
				},
			},
		},
		{
			name: "January issues",
			config: Config{
				Defaults: defaults,
				Issues:   []Issue{issue1, issue1_3, issue2_4},
			},
			month: Month(1),
			issuesToCreate: IssuesToCreate{
				Issues: []IssueToCreate{
					{
						Issue:      issue1,
						ProjectID:  "default_project_id",
						TargetRepo: "default/repo",
					},
					{
						Issue:      issue1_3,
						ProjectID:  "default_project_id",
						TargetRepo: "default/repo",
					},
				},
			},
		},
		{
			name: "February issues",
			config: Config{
				Defaults: defaults,
				Issues:   []Issue{issue2, issue1_3, issue2_4},
			},
			month: Month(2),
			issuesToCreate: IssuesToCreate{
				Issues: []IssueToCreate{
					{
						Issue:      issue2,
						ProjectID:  "default_project_id",
						TargetRepo: "default/repo",
					},
					{
						Issue:      issue2_4,
						ProjectID:  "default_project_id",
						TargetRepo: "default/repo",
					},
				},
			},
		},
		{
			name: "Override defaults",
			config: Config{
				Defaults: defaults,
				Issues:   []Issue{issue_project_repo},
			},
			month: Month(1),
			issuesToCreate: IssuesToCreate{
				Issues: []IssueToCreate{
					{
						Issue:      issue_project_repo,
						ProjectID:  "other_project_id",
						TargetRepo: "other/repo",
					},
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIssuesToCreate(tt.config, tt.month)
			if !got.Equals(tt.issuesToCreate) {
				t.Errorf("expected %v, got %v", tt.issuesToCreate, got)
			}
		})
	}
}
