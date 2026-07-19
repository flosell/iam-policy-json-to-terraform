package converter

import (
	"regexp"
	"sort"
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
	for _, k := range sortedKeys(conditions) {
		v := conditions[k]
		for _, k2 := range sortedKeys(v) {
			result = append(result, hclCondition{
				Test:     k,
				Variable: k2,
				Values:   convertStringOrStringArray(v[k2]),
			})
		}
	}
	return result
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
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
		// revive:disable:unchecked-type-assertion
		m := v.(map[string]interface{})
		result := make([]hclPrincipal, 0)
		for _, k := range sortedKeys(m) {
			result = append(result, hclPrincipal{
				Type:        k,
				Identifiers: convertStringOrStringArray(m[k]),
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
