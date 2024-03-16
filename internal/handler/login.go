package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/controller"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Função para lidar com o login do usuário
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decodifica o JSON da solicitação para a struct loginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Busca o usuário pelo e-mail fornecido
	user := controller.FindUserByEmail(loginData.Email)
	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	// Verifica se a senha fornecida está correta
	err = controller.VerifyPassword(user.Password, loginData.Password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Se as credenciais estiverem corretas, gera um token JWT para o usuário
	token, err := controller.GenerateJWT(*user)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	// Retorna o token JWT para o usuário
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}
