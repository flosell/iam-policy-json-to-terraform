package converter

import "encoding/json"

type jsonPolicyDocument struct {
	Version string
	Statement []jsonStatement
}

type stringOrStringArray interface{}

type jsonStatement struct {
	Sid         string
	Effect      string
	Resource    stringOrStringArray
	NotResource stringOrStringArray
	Action      stringOrStringArray
	NotAction   stringOrStringArray
	Condition   map[string]map[string]string
}

func decode(b []byte) ([]jsonStatement, error) {
	document := &jsonPolicyDocument{}
	err := json.Unmarshal(b, document)

	if err != nil {
		return nil, err
	}

	return document.Statement, nil

}
