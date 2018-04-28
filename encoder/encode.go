package encoder

import "github.com/flosell/hclencoder"

func Encode(dataSource DataSource) (string, error) {
	b, err :=  hclencoder.Encode(WholeFile{Entry: dataSource})

	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
