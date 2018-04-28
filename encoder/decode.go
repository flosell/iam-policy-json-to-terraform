package encoder

import "encoding/json"

type JSONPolicyDocument struct {
	Version string
	Statement []Statement
}

func Decode(b []byte) ([]Statement, error) {
	document := &JSONPolicyDocument{}
	err := json.Unmarshal(b, document)

	if err != nil {
		return nil, err
	}

	return document.Statement, nil

}
