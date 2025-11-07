package main

type Config struct {
	Issues []Issue `yaml:"issues"`
}

type Issue struct {
	Name           string `yaml:"name"`
	CreationMonths []int  `yaml:"creation_months"`
}
