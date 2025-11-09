package main

import (
	"context"
	"flag"
	"log"
	"os"
)

func main() {
	var (
		month      = flag.Int("month", 0, "Month (1-12) to filter issues")
		configFile = flag.String("config", "", "Path to config file (required)")
		debug      = flag.Bool("debug", false, "Enable debug logging")
	)
	flag.Parse()

	SetDebugMode(*debug)

	monthEnum, err := ParseMonth(*month)
	if err != nil {
		log.Fatalf("failed to parse month: %v", err)
	}

	configPath := *configFile
	if configPath == "" {
		log.Fatalf("config file is required. Use --config to specify a config file")
	}

	config, err := LoadConfig(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("config file not found: %s\nUse --config to specify a different config file", configPath)
		}
		log.Fatalf("failed to load config: %v", err)
	}

	// Create GitHub client for validation
	ghClient, err := NewGitHubClient()
	if err != nil {
		log.Fatalf("failed to create GitHub client: %v", err)
	}

	if err := ValidateConfig(config, ghClient); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	issuesToCreate := GetIssuesToCreate(config, monthEnum)

	ctx := context.Background()
	if err := outputJSON(ctx, issuesToCreate, config.Defaults, ghClient); err != nil {
		log.Fatalf("failed to output JSON: %v", err)
	}
}

