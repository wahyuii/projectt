package schema

import (
	usersSchema "projectt/app/graphql/schema/users"

	"github.com/graphql-go/graphql"
)

func initRootMutation() graphql.ObjectConfig {
	mutationFields := graphql.Fields{
		"insert_hi": usersSchema.InsertHISchema(),
	}
	return graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	}
}
