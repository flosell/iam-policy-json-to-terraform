package converter

func convertStatements(json JSONStatement) HCLStatement {
	return HCLStatement{
		Effect: json.Effect,
		Sid:json.Sid,
		Resources:[]string {json.Resource},
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
