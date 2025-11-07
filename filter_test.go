package main

import (
	"testing"
)

func TestGetIssuesToCreate(t *testing.T) {
	issue1 := Issue{
		Name:           "Issue 1",
		CreationMonths: []int{1},
	}
	issue2 := Issue{
		Name:           "Issue 2",
		CreationMonths: []int{2},
	}
	issue1_3 := Issue{
		Name:           "Issue 1_3",
		CreationMonths: []int{1, 3},
	}
	issue2_4 := Issue{
		Name:           "Issue 2_4",
		CreationMonths: []int{2, 4},
	}

	cases := []struct {
		name           string
		issues         []Issue
		month          int
		issuesToCreate IssuesToCreate
	}{
		{
			name:           "No issues",
			issues:         []Issue{},
			month:          1,
			issuesToCreate: IssuesToCreate{},
		},
		{
			name:   "One issue",
			issues: []Issue{issue1},
			month:  1,
			issuesToCreate: IssuesToCreate{
				Issues: []Issue{issue1},
			},
		},
		{
			name:   "January issues",
			issues: []Issue{issue1, issue1_3, issue2_4},
			month:  1,
			issuesToCreate: IssuesToCreate{
				Issues: []Issue{issue1, issue1_3},
			},
		},
		{
			name:   "February issues",
			issues: []Issue{issue2, issue1_3, issue2_4},
			month:  2,
			issuesToCreate: IssuesToCreate{
				Issues: []Issue{issue2, issue2_4},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIssuesToCreate(tt.issues, tt.month)
			if !got.Equals(tt.issuesToCreate) {
				t.Errorf("expected %v, got %v", tt.issuesToCreate, got)
			}
		})
	}
}
