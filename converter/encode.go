package converter

import "github.com/flosell/hclencoder"

type hclCondition struct {
	Test     string   `hcl:"test"`
	Variable string   `hcl:"variable"`
	Values   []string `hcl:"values"`
}

type hclStatement struct {
	Sid          string         `hcl:"sid"`
	Effect       string         `hcl:"effect"`
	Resources    []string       `hcl:"resources"`
	NotResources []string       `hcl:"not_resources"`
	Actions      []string       `hcl:"actions"`
	NotActions   []string       `hcl:"not_actions"`
	Conditions   []hclCondition `hcl:"condition,squash"`
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
