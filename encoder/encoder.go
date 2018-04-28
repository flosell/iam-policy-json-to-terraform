package encoder

import (
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
)

type Statement struct {
	Sid string
}

type StatementList struct {
	Statements []*Statement
}

type PolicyDocument struct {
	Name       string
	Statements StatementList
}

func (s *Statement) Encode() *ast.ObjectItem {
	content := &ast.ObjectList{
		Items: []*ast.ObjectItem{},
	}

	content.Add(&ast.ObjectItem{
		Keys: []*ast.ObjectKey{
			&ast.ObjectKey{Token: token.Token{Type: token.IDENT, Text: `sid`}},
			&ast.ObjectKey{Token: token.Token{Type: token.ASSIGN, Text: `=`}},
			&ast.ObjectKey{Token: token.Token{Type: token.STRING, Text: s.Sid}},
		},
	})

	return &ast.ObjectItem{
		Keys: []*ast.ObjectKey{
			&ast.ObjectKey{Token: token.Token{Type: token.STRING, Text: `statement`}},
		},
		Val: &ast.ObjectType{
			List: content,
		},
	}
}

func (sl *StatementList) Encode() *ast.ObjectList {
	objectList := make([]*ast.ObjectItem,len(sl.Statements))
	for i, s := range sl.Statements {
		objectList[i] = s.Encode()
	}
	return &ast.ObjectList{
		Items: objectList,
	}
}

func (p *PolicyDocument) Encode() *ast.ObjectList {
	return &ast.ObjectList{
		Items: []*ast.ObjectItem{
			&ast.ObjectItem{
				Keys: []*ast.ObjectKey{
					&ast.ObjectKey{Token: token.Token{Type: token.STRING, Text: `data`}},
					&ast.ObjectKey{Token: token.Token{Type: token.STRING, Text: `"aws_iam_policy_document"`}},
					&ast.ObjectKey{Token: token.Token{Type: token.STRING, Text: `"` + p.Name + `"`}},
				},
				Val: &ast.ObjectType{
					List: p.Statements.Encode(),
				},
			},
		},
	}
}
