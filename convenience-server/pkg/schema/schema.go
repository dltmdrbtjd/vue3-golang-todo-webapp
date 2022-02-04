package schema

import (
	"io/ioutil"

	"github.com/Convenience-Tools/convenience-server/pkg/resolver"
	"github.com/graph-gophers/graphql-go"
)

var Schema *graphql.Schema

func init() {
	s, err := getSchema("./schema/schema.graphql")
	if err != nil {
		panic(err)
	}

	graphql.MustParseSchema(s, &resolver.RootResolver{}, graphql.UseStringDescriptions())
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
