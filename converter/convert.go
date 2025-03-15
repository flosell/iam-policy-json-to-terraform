package converter

import (
	"regexp"
	"strconv"
)

func escapePolicyVariables(s string) string {
	// Escape TF special characters
	re := regexp.MustCompile(
		`\${(` +
			`([^}]*:[^}]*)` +
			`|` +
			`([*?$])` +
			`)}`)
	return re.ReplaceAllString(s, "$$$${$1}")
}

func convertConditions(conditions map[string]map[string]stringOrStringArray) []hclCondition {
	result := make([]hclCondition, 0)
	for k, v := range conditions {
		for k2, v2 := range v {
			result = append(result, hclCondition{
				Test:     k,
				Variable: k2,
				Values:   convertStringOrStringArray(v2),
			})
		}
	}
	return result
}

func convertPrincipals(v stringOrMapWithStringOrStringArray) []hclPrincipal {
	switch v.(type) {
	case string:
		return []hclPrincipal{
			{
				Type:        "*",
				Identifiers: []string{v.(string)},
			},
		}
	case map[string]interface{}:
		result := make([]hclPrincipal, 0)
		// revive:disable:unchecked-type-assertion
		for k, v := range v.(map[string]interface{}) {
			result = append(result, hclPrincipal{
				Type:        k,
				Identifiers: convertStringOrStringArray(v),
			})
		}
		return result
	default:
		return nil
	}
}

func convertStringOrStringArray(v stringOrStringArray) []string {
	switch v.(type) {
	case []interface{}:
		var resources []string
		resourceArray, _ := v.([]interface{}) // revive:disable:unchecked-type-assertion
		resources = make([]string, len(resourceArray))
		for i, r := range resourceArray {
			resources[i] = escapePolicyVariables(r.(string))
		}
		return resources
	case string:
		return []string{v.(string)}
	case bool:
		return []string{strconv.FormatBool(v.(bool))}
	default:
		return nil
	}
}

func convertStatements(json jsonStatement) hclStatement {
	return hclStatement{
		Effect:        json.Effect,
		Sid:           json.Sid,
		Resources:     convertStringOrStringArray(json.Resource),
		NotResources:  convertStringOrStringArray(json.NotResource),
		Actions:       convertStringOrStringArray(json.Action),
		NotActions:    convertStringOrStringArray(json.NotAction),
		Conditions:    convertConditions(json.Condition),
		Principals:    convertPrincipals(json.Principal),
		NotPrincipals: convertPrincipals(json.NotPrincipal),
	}
}

// Convert reads a JSON policy document and return a string with a terraform policy document definition
func Convert(policyName string, b []byte) (string, error) {
	statementsFromJSON, err := decode(b)

	if err != nil {
		return "", err
	}

	hclStatements := make([]hclStatement, len(statementsFromJSON))

	for i, s := range statementsFromJSON {
		hclStatements[i] = convertStatements(s)
	}

	dataSource := hclDataSource{
		Type:       "aws_iam_policy_document",
		Name:       policyName,
		Statements: hclStatements,
	}

	tfFromStatements, err := encode(dataSource)

	if err != nil {
		return "", err
	}

	return tfFromStatements, err
}
