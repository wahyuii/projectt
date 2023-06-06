package usersRepository

import (
	usersModels "projectt/app/database/models/users"

	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	tapDB *sqlx.DB
	db    *sqlx.DB
}

type UsersRepository interface {
	GetUsersRepo() (*usersModels.User, error)
	InsertHIRepo() (usersModels.User, error)
}

func NewUsersRepository(tapDB *sqlx.DB) UsersRepository {
	return &usersRepository{
		tapDB: tapDB,
	}
}

func (UsersRepository *usersRepository) InsertHIRepo() (usersModels.User, error) {
	query := "INSERT INTO users (name, email, age) VALUES ($1, $2, $3)"
	user := usersModels.User{}
	var userID int64
	err := UsersRepository.db.QueryRow(query, user.Name, user.Email, user.Age).Scan(&userID)
	if err != nil {
		return usersModels.User{}, err
	}

	return user, nil
}

func (cekRepo *usersRepository) GetUsersRepo() (*usersModels.User, error) {
	result := &usersModels.User{}
	rows := cekRepo.tapDB.QueryRow(
		`SELECT
			id, 
			name,
			email
        FROM users
        `,
	)
	sqlError := rows.Scan(
		&result.ID,
		&result.Name,
		&result.Email,
	)

	return result, sqlError
}
