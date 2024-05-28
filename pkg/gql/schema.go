package gql

import "github.com/graphql-go/graphql"

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})

func executeQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			println("err:", err.Message)
		}
	}

	return result
}
