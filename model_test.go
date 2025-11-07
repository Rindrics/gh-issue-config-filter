package main

import (
	"testing"
)

func TestMonth(t *testing.T) {
	cases := []struct {
		name        string
		digit       int
		expectError bool
	}{
		{name: "January", digit: 1, expectError: false},
		{name: "February", digit: 2, expectError: false},
		{name: "December", digit: 12, expectError: false},
		{name: "Invalid", digit: 0, expectError: true},
		{name: "Invalid", digit: 13, expectError: true},
		{name: "Invalid", digit: -1, expectError: true},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseMonth(tt.digit)
			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestIssue_IsCreationMonth(t *testing.T) {

	cases := []struct {
		name           string
		creationMonths []Month
		thisMonth      Month
		expect         bool
	}{
		{
			name:           "Matches first month",
			creationMonths: []Month{1, 2, 3},
			thisMonth:      1,
			expect:         true,
		},
		{
			name:           "Matches second month",
			creationMonths: []Month{1, 2, 3},
			thisMonth:      2,
			expect:         true,
		},
		{
			name:           "Does not match",
			creationMonths: []Month{1, 2, 3},
			thisMonth:      4,
			expect:         false,
		},
		{
			name:           "Empty creation months",
			creationMonths: []Month{},
			thisMonth:      3,
			expect:         false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			issue := Issue{
				CreationMonths: tt.creationMonths,
			}
			if issue.IsCreationMonth(tt.thisMonth) != tt.expect {
				t.Errorf("expected %v, got %v", tt.expect, issue.IsCreationMonth(tt.thisMonth))
			}
		})
	}
}
