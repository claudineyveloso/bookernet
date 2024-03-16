package controller

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Chave secreta usada para assinar tokens JWT. Mantenha isso seguro em um ambiente de produção.
var jwtKey = []byte("chave_secreta_do_jwt")

// Struct para armazenar informações do usuário
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []User{
	{
		Email:    "claudineyveloso@gmail.com",
		Password: "$2a$10$Upds6QH3WRt9IR8dxA18eOTHrG2dKoLJgcTao6syUR4KTb6XE1k.2",
	},
}

// Função para verificar se a senha fornecida é válida
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Função para gerar um token JWT para o usuário
func GenerateJWT(user User) (string, error) {
	// Define os claims do token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expira em 1 dia
	}

	// Cria o token JWT com os claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta e retorna o token como uma string
	return token.SignedString(jwtKey)
}

func FindUserByEmail(email string) *User {
	for _, user := range users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
