package main

import (
	"reflect"
	"slices"
)

type Config struct {
	Issues []Issue `yaml:"issues"`
}

type Issue struct {
	Name           string `yaml:"name"`
	CreationMonths []int  `yaml:"creation_months"`
}

type IssueToCreate struct {
	Issue  Issue
	Fields map[string]string
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
