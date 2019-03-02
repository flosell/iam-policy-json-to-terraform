package converter

import "encoding/json"

type jsonPolicyDocument struct {
	Version   string
	Statement []jsonStatement
}

type stringOrStringArray interface{}
type stringOrMapWithStringOrStringArray interface{}

type jsonStatement struct {
	Sid          string
	Effect       string
	Resource     stringOrStringArray
	NotResource  stringOrStringArray
	Action       stringOrStringArray
	NotAction    stringOrStringArray
	Condition    map[string]map[string]stringOrStringArray
	Principal    stringOrMapWithStringOrStringArray
	NotPrincipal stringOrMapWithStringOrStringArray
}

func decode(b []byte) ([]jsonStatement, error) {
	document := &jsonPolicyDocument{}
	err := json.Unmarshal(b, document)

	if err != nil {
		return nil, err
	}

	return document.Statement, nil

}
