package converter

import (
	"strings"
)

func convertStrings(s string) string {
	// Escape TF special characters
	return strings.Replace(s, "$", "$$", -1)
}

func convertConditions(conditions map[string]map[string]string) []HCLCondition {
	result := make([]HCLCondition, 0)
	for k, v := range conditions {
		for k2, v2 := range v {
			result = append(result, HCLCondition{
				Test:     k,
				Variable: k2,
				Values:   []string{convertStrings(v2)},
			})
		}
	}
	return result
}

func convertStringOrStringArray(v StringOrStringArray) []string {
	if v == nil {
		return nil
	}

	resourceArray, arrOk := v.([]interface{})
	resourceString, _ := v.(string)
	var resources []string
	if arrOk {
		resources = make([]string, len(resourceArray))
		for i, r := range resourceArray {
			resources[i] = convertStrings(r.(string))
		}
	} else {
		resources = []string{resourceString}
	}
	return resources
}

func convertStatements(json JSONStatement) HCLStatement {
	return HCLStatement{
		Effect:     json.Effect,
		Sid:        json.Sid,
		Resources:  convertStringOrStringArray(json.Resource),
		NotResources:  convertStringOrStringArray(json.NotResource),
		Actions:    convertStringOrStringArray(json.Action),
		NotActions: convertStringOrStringArray(json.NotAction),
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
