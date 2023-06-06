package usersSchema

import (
	"projectt/app/context"
	usersResolver "projectt/app/graphql/resolver/users"
	usersTypedefs "projectt/app/graphql/type-def/users"

	"github.com/graphql-go/graphql"
)

func ViewUsersSchema() *graphql.Field {
	return &graphql.Field{
		Type:        usersTypedefs.GetUsers(),
		Description: "Mengambil data di database Cek",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			DefineCekResolver := params.Context.Value(context.UsersResolverKey).(usersResolver.UsersResolver)
			return DefineCekResolver.GetUsersResolver(params)
		},
	}
}
func InsertHISchema() *graphql.Field {
	return &graphql.Field{
		Type:        usersTypedefs.GetUsers(),
		Description: "Insert Hi",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"age": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			DefineUsersResolver := params.Context.Value(context.UsersResolverKey).(usersResolver.UsersResolver)
			return DefineUsersResolver.InsertHIResolver(params)

		},
	}
}
