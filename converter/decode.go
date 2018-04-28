package converter

import "encoding/json"

type JSONPolicyDocument struct {
	Version string
	Statement []HclStatement
}

func Decode(b []byte) ([]HclStatement, error) {
	document := &JSONPolicyDocument{}
	err := json.Unmarshal(b, document)

	if err != nil {
		return nil, err
	}

	return document.Statement, nil

}
