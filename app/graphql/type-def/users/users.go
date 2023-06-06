package usersTypedefs

import "github.com/graphql-go/graphql"

func GetUsers() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "users",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"name":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
			"age":   &graphql.Field{Type: graphql.Int},
		},
	})
}
