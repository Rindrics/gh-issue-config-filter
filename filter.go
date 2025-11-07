package main

func GetIssuesToCreate(candidates []Issue, month int) IssuesToCreate {
	issuesToCreate := IssuesToCreate{
		Issues: []Issue{},
	}

	for _, candidate := range candidates {
		if candidate.IsCreationMonth(month) {
			issuesToCreate.Issues = append(issuesToCreate.Issues, candidate)
		}
	}
	return issuesToCreate
}
