package main

import "errors"

func ValidateIssue(candidate IssueToCreate) error {
	if candidate.Issue.Name == "" {
		return errors.New("issue name is required")
	}

	return nil
}
