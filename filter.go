package main

func GetIssuesToCreate(candidates []Issue, month int) IssuesToCreate {
	issuesToCreate := IssuesToCreate{
		Issues: []IssueToCreate{},
	}

	for _, candidate := range candidates {
		if candidate.IsCreationMonth(month) {
			issuesToCreate.Issues = append(issuesToCreate.Issues, IssueToCreate{Issue: candidate})
		}
	}
	return issuesToCreate
}
