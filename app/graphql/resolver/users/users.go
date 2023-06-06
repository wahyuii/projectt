package users

import (
	usersModels "projectt/app/database/models/users"
	"projectt/app/database/repository"
	usersRepository "projectt/app/database/repository/users"

	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

type usersResolver struct {
	tapDB           *sqlx.DB
	transact        repository.Transact
	usersRepository usersRepository.UsersRepository
}

type UsersResolver interface {
	GetUsersResolver(params graphql.ResolveParams) (*usersModels.User, error)
	InsertHIResolver(params graphql.ResolveParams) (*usersModels.User, error)
}

func NewUsersResolver(
	_tapDB *sqlx.DB,
	_transact repository.Transact,
	_usersRepository usersRepository.UsersRepository,
) UsersResolver {
	return &usersResolver{
		tapDB:           _tapDB,
		transact:        _transact,
		usersRepository: _usersRepository,
	}
}

func (usersResolver *usersResolver) GetUsersResolver(params graphql.ResolveParams) (*usersModels.User, error) {

	rawData, err := usersResolver.usersRepository.GetUsersRepo()

	return rawData, err
}

func (usersResolver *usersResolver) InsertHIResolver(params graphql.ResolveParams) (*usersModels.User, error) {

	name, _ := params.Args["name"].(string)
	email, _ := params.Args["email"].(string)
	age, _ := params.Args["age"].(int)
	user := &usersModels.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	_, err := usersResolver.usersRepository.InsertHIRepo()

	return user, err
}
