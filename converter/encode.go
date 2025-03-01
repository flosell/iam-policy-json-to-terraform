package converter

import "github.com/flosell/hclencoder"

type hclCondition struct {
	Test     string   `hcl:"test"`
	Variable string   `hcl:"variable"`
	Values   []string `hcl:"values"`
}

type hclPrincipal struct {
	Type        string   `hcl:"type"`
	Identifiers []string `hcl:"identifiers"`
}

type hclStatement struct {
	Sid           string         `hcl:"sid" hcle:"omitempty"`
	Effect        string         `hcl:"effect"`
	Resources     []string       `hcl:"resources"`
	NotResources  []string       `hcl:"not_resources"`
	Actions       []string       `hcl:"actions"`
	NotActions    []string       `hcl:"not_actions"`
	Conditions    []hclCondition `hcl:"condition,squash"`
	Principals    []hclPrincipal `hcl:"principals,squash"`
	NotPrincipals []hclPrincipal `hcl:"not_principals,squash"`
}

type hclDataSource struct {
	Type       string         `hcl:",key"`
	Name       string         `hcl:",key"`
	Statements []hclStatement `hcl:"statement,squash"`
}

type hclWholeFile struct {
	Entry hclDataSource `hcl:"data"`
}

func encode(dataSource hclDataSource) (string, error) {
	b, err := hclencoder.Encode(hclWholeFile{Entry: dataSource})

	if err != nil {
		return "", err
	}
	return string(b), nil
}
