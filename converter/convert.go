package converter

func Convert(b []byte) (string, error) {
	statementsFromJson, err := Decode(b)

	if err != nil {
		return "", err
	}

	data_source := HclDataSource{
		Type: "aws_iam_policy_document",
		Name: "deny_access_without_mfa",
		Statements: statementsFromJson,
	}

	tfFromStatements, err := Encode(data_source)

	if err != nil {
		return "", err
	}

	return tfFromStatements, err

}
