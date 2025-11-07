package main

func GetIssuesToCreate(config Config, month Month) IssuesToCreate {
	issuesToCreate := IssuesToCreate{
		Issues: []IssueToCreate{},
	}

	for _, candidate := range config.Issues {
		if candidate.IsCreationMonth(month) {
			issueToCreate := IssueToCreate{Issue: candidate}
			if candidate.ProjectID != "" {
				issueToCreate.ProjectID = candidate.ProjectID
			} else {
				issueToCreate.ProjectID = config.Defaults.ProjectID
			}
			if candidate.TargetRepo != "" {
				issueToCreate.TargetRepo = candidate.TargetRepo
			} else {
				issueToCreate.TargetRepo = config.Defaults.TargetRepo
			}
			issuesToCreate.Issues = append(issuesToCreate.Issues, issueToCreate)
		}
	}
	return issuesToCreate
}
