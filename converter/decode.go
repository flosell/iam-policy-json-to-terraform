package converter

import (
	"encoding/json"
)

type jsonStatements []jsonStatement
type jsonPolicyDocument struct {
	Version   string
	Statement jsonStatements
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

func (stmts *jsonStatements) UnmarshalJSON(b []byte) error {
	var jsonStatements []jsonStatement
	errSliceUnmarshal := json.Unmarshal(b, &jsonStatements)
	if errSliceUnmarshal == nil {
		*stmts = jsonStatements
		return nil
	}

	if e, ok := errSliceUnmarshal.(*json.UnmarshalTypeError); ok {
		if e.Value == "object" {
			var s jsonStatement
			errStringUnmarshal := json.Unmarshal(b, &s)
			if errStringUnmarshal != nil {
				return errStringUnmarshal
			}
			*stmts = []jsonStatement{s}
			return nil
		}
	}
	return errSliceUnmarshal
}

func decode(b []byte) ([]jsonStatement, error) {
	document := &jsonPolicyDocument{}
	err := json.Unmarshal(b, document)

	if err != nil {
		return nil, err
	}

	return document.Statement, nil

}
