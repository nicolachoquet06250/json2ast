package json2ast

import "github.com/nicolachoquet06250/json2ast/dora"

func ToAst(json string) (*dora.Client, error) {
	c, err := dora.NewFromString(json)
	if err != nil {
		return nil, err
	}

	return c, nil
}
