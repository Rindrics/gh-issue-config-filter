package main

import (
	"reflect"
	"slices"
)

type Defaults struct {
	ProjectID  string `yaml:"project_id"`
	TargetRepo string `yaml:"target_repo"`
}

type Config struct {
	Defaults Defaults `yaml:"defaults"`
	Issues   []Issue  `yaml:"issues"`
}

type Issue struct {
	Name           string `yaml:"name"`
	CreationMonths []int  `yaml:"creation_months"`
	ProjectID      string `yaml:"issue.project_id"`
	TargetRepo     string `yaml:"issue.target_repo"`
}

type IssueToCreate struct {
	Issue      Issue
	Fields     map[string]string
	ProjectID  string
	TargetRepo string
}

type IssuesToCreate struct {
	Issues []IssueToCreate
}

func (i *IssuesToCreate) Equals(other IssuesToCreate) bool {
	if len(i.Issues) == 0 && len(other.Issues) == 0 {
		return true
	}
	return reflect.DeepEqual(i.Issues, other.Issues)
}

func (i *Issue) IsCreationMonth(month int) bool {
	return slices.Contains(i.CreationMonths, month)
}
