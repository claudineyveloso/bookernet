package user

import (
	"context"
	"database/sql"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUsers() ([]*types.User, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbUsers, err := queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []*types.User
	for _, dbUser := range dbUsers {
		// user := &types.User{
		// 	ID:        dbUser.ID,
		// 	Email:     dbUser.Email,
		// 	Password:  dbUser.Password,
		// 	IsActive:  dbUser.IsActive,
		// 	UserType:  dbUser.UserType,
		// 	CreatedAt: dbUser.CreatedAt,
		// 	UpdatedAt: dbUser.UpdatedAt,
		// }
		user := convertDBUserToUser(dbUser)
		users = append(users, user)
	}
	return users, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func convertDBUserToUser(dbUser db.User) *types.User {
	user := &types.User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		IsActive:  dbUser.IsActive,
		UserType:  dbUser.UserType,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	return user
}
