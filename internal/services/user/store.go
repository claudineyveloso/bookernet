package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.CreateUserPayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	user.ID = uuid.New()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao gerar hash da senha:", err)
		return err
	}
	hashedPasswordString := string(hashedPassword)

	createUserParams := db.CreateUserParams{
		ID:        user.ID,
		Email:     user.Email,
		Password:  hashedPasswordString,
		IsActive:  user.IsActive,
		UserType:  user.UserType,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	//ctx := context.Background()

	if err := queries.CreateUser(ctx, createUserParams); err != nil {
		//http.Error(_, "Erro ao criar usuário", http.StatusInternalServerError)
		return err
	}
	return nil
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
