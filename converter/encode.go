package converter

import "github.com/flosell/hclencoder"

type HCLStatement struct {
	Sid    string `hcl:"sid"`
	Effect string `hcl:"effect"`
	Resources []string `hcl:"resources"`
	NotActions []string `hcl:"not_actions"`
}

type HCLDataSource struct {
	Type       string         `hcl:",key"`
	Name       string         `hcl:",key"`
	Statements []HCLStatement `hcl:"statement,squash"`
}

type HclWholeFile struct {
	Entry HCLDataSource `hcl:"data"`
}

func Encode(dataSource HCLDataSource) (string, error) {
	b, err := hclencoder.Encode(HclWholeFile{Entry: dataSource})

	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
