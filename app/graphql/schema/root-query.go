package schema

import (
	usersSchema "projectt/app/graphql/schema/users"

	"github.com/graphql-go/graphql"
)

func initRootQuery() graphql.ObjectConfig {
	queryFields := graphql.Fields{
		"view_cek": usersSchema.ViewUsersSchema(),
	}
	return graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	}
}
