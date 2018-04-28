package converter

import "encoding/json"

type JSONPolicyDocument struct {
	Version string
	Statement []JSONStatement
}

type StringOrStringArray interface{}

type JSONStatement struct {
	Sid    string
	Effect string
	Resource StringOrStringArray
	NotResource StringOrStringArray
	Action StringOrStringArray
	NotAction StringOrStringArray
	Condition map[string]map[string]string
}

func Decode(b []byte) ([]JSONStatement, error) {
	document := &JSONPolicyDocument{}
	err := json.Unmarshal(b, document)

	if err != nil {
		return nil, err
	}

	return document.Statement, nil

}
