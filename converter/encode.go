package converter

import "github.com/flosell/hclencoder"

type HclStatement struct {
	Sid    string `hcl:"sid"`
	Effect string `hcl:"effect"`
}

type HclDataSource struct {
	Type       string         `hcl:",key"`
	Name       string         `hcl:",key"`
	Statements []HclStatement `hcl:"statement,squash"`
}

type HclWholeFile struct {
	Entry HclDataSource `hcl:"data"`
}

func Encode(dataSource HclDataSource) (string, error) {
	b, err := hclencoder.Encode(HclWholeFile{Entry: dataSource})

	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
