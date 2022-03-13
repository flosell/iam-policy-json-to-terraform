package converter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type jsonStatements []jsonStatement
type jsonPolicyDocument struct {
	Version        string
	Statement      *jsonStatements
	PolicyName     *interface{}
	PolicyDocument *interface{}
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

// ErrorLackOfStatements indicates that the input JSON did not contain any statements, so likely isn't a useful input
var ErrorLackOfStatements = errors.New("input did not contain any statements")

// ErrorLooksLikeCloudformation indicates that the input JSON did not contain any statements but looks like CloudFormation
var ErrorLooksLikeCloudformation = errors.New("input did not contain any statements " +
	"but looks like CloudFormation code - this is currently not supported. Look at GitHub Issue #50 for details")

func decode(b []byte) ([]jsonStatement, error) {
	document := &jsonPolicyDocument{}
	err := json.Unmarshal(escapeHclSnippetsInJSON(b), document)

	if err != nil {
		return nil, err
	}

	if document.Statement == nil {
		if document.PolicyName != nil || document.PolicyDocument != nil {
			return nil, fmt.Errorf("invalid policy input: %w", ErrorLooksLikeCloudformation)
		}
		return nil, fmt.Errorf("invalid policy input: %w", ErrorLackOfStatements)
	}

	return *document.Statement, nil

}

func escapeHclSnippetsInJSON(b []byte) []byte {
	unescapedBuffer := bytes.Buffer{}
	escapeBuffer := bytes.Buffer{}

	currentBuffer := &unescapedBuffer

	expressionDepth := 0

	for i := 0; i < len(b); i++ {
		if expressionDepth == 0 {
			unescapedBuffer.WriteByte(b[i])
		} else {
			if b[i] == '}' {
				expressionDepth--
				if expressionDepth == 0 {
					currentBuffer = &unescapedBuffer
					currentBuffer.Write(escapeJSON(escapeBuffer.Bytes()))
					escapeBuffer.Reset()
					currentBuffer.WriteByte(b[i])
				} else {
					currentBuffer.WriteByte(b[i])
				}
			} else {
				currentBuffer.WriteByte(b[i])
			}
		}

		if b[i] == '$' && len(b) > i+1 && b[i+1] == '{' {
			currentBuffer.WriteByte(b[i+1])
			expressionDepth++
			i++
			currentBuffer = &escapeBuffer
		}
	}
	unescapedBuffer.Write(escapeBuffer.Bytes())
	return unescapedBuffer.Bytes()
}

func escapeJSON(b []byte) []byte {
	marshalled, _ := json.Marshal(string(b))
	withoutQuotes := marshalled[1 : len(marshalled)-1]
	return withoutQuotes
}
