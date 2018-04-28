package converter

func convertConditions(conditions map[string]map[string]string) []HCLCondition {
	result := make([]HCLCondition, 0)
	for k, v := range conditions {
		for k2, v2 := range v {
			result = append(result, HCLCondition{
				Test:     k,
				Variable: k2,
				Values:   []string{v2},
			})
		}
	}
	return result
}

func convertStatements(json JSONStatement) HCLStatement {
	return HCLStatement{
		Effect: json.Effect,
		Sid:json.Sid,
		Resources:[]string {json.Resource},
		Actions: json.Action,
		NotActions: json.NotAction,
		Conditions: convertConditions(json.Condition),
	}
}

func Convert(b []byte) (string, error) {
	statementsFromJson, err := Decode(b)
	hclStatements := make([]HCLStatement , len(statementsFromJson))

	for i, s := range statementsFromJson {
		hclStatements[i] = convertStatements(s)
	}

	if err != nil {
		return "", err
	}

	data_source := HCLDataSource{
		Type: "aws_iam_policy_document",
		Name: "deny_access_without_mfa",
		Statements: hclStatements,
	}

	tfFromStatements, err := Encode(data_source)

	if err != nil {
		return "", err
	}

	return tfFromStatements, err

}
