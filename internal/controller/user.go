package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserRequest struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	IsActive  bool      `json:"is_active"`
	UserType  string    `json:"user_type" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUserController(rw http.ResponseWriter, r *http.Request, queries *db.Queries) {

	var createUserRequest UserRequest

	log.Println("Corpo da solicitação:", r.Body)

	if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
		http.Error(rw, "Erro ao decodificar corpo da solicitação"+err.Error(), http.StatusBadRequest)
		return
	}

	// Validar os campos do usuário antes de continuar
	if err := createUserRequest.Validate(); err != nil {
		http.Error(rw, "Erro de validação: "+err.Error(), http.StatusBadRequest)
		return
	}

	createUserRequest.ID = uuid.New()
	now := time.Now()
	createUserRequest.CreatedAt = now
	createUserRequest.UpdatedAt = now

	// Create instance of db.CreateUserParams based on request data
	createUserParams := db.CreateUserParams{
		ID:        createUserRequest.ID,
		Email:     createUserRequest.Email,
		Password:  createUserRequest.Password,
		IsActive:  createUserRequest.IsActive,
		UserType:  createUserRequest.UserType,
		CreatedAt: createUserRequest.CreatedAt,
		UpdatedAt: createUserRequest.UpdatedAt,
	}

	if err := queries.CreateUser(r.Context(), createUserParams); err != nil {
		http.Error(rw, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	// Encode user data in JSON format and send it as a response
	if err := json.NewEncoder(rw).Encode(createUserParams); err != nil {
		http.Error(rw, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Email":
				errorMap["email"] = "O campo de e-mail não pode estar vazio"
			case "Password":
				errorMap["password"] = "O campo de senha não pode estar vazio"
			case "UserType":
				errorMap["user_type"] = "O campo de tipo de usuário não pode estar vazio"
				// Adicionar casos para outros campos, se necessário
			}
		}
		// Construir mensagem de erro concatenando as mensagens de erro específicas
		var errorMsg string
		for _, msg := range errorMap {
			errorMsg += msg + "\n"
		}

		return errors.New(errorMsg)
	}
	return nil
}

func GetUsersController(rw http.ResponseWriter, r *http.Request, queries *db.Queries) {
	users, err := queries.GetUsers(r.Context())
	if err != nil {
		http.Error(rw, fmt.Sprintf("Erro ao obter usuário: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		http.Error(rw, "Erro ao codificar em JSON", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	_, err = rw.Write(jsonData)
	if err != nil {
		http.Error(
			rw,
			"Erro ao escrever resposta",
			http.StatusInternalServerError,
		)
		return
	}

}

func GetUserController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	userID := r.URL.Query().Get("id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "ID do usuário não pode ser vazio", http.StatusBadRequest)
		return
	}
	user, err := queries.GetUser(r.Context(), parsedUserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter usuário: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		http.Error(w, "Erro ao codificar em JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
		return
	}
}

func UpdateUserController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var updateUserRequest UserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateUserRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação: "+err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now()

	// Create instance of db.UpdateUserParams based on request data
	updateUserParams := db.UpdateUserParams{
		ID:        updateUserRequest.ID,
		Email:     updateUserRequest.Email,
		IsActive:  updateUserRequest.IsActive,
		UserType:  updateUserRequest.UserType,
		UpdatedAt: now,
	}

	if err := queries.UpdateUser(r.Context(), updateUserParams); err != nil {
		http.Error(w, "Erro ao alterar dados do usuário", http.StatusInternalServerError)
		return
	}

	// Encode user data in JSON format and send it as a response
	if err := json.NewEncoder(w).Encode(updateUserParams); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func UpdatePasswordController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var updatePasswordUserRequest UserRequest

	if err := json.NewDecoder(r.Body).Decode(&updatePasswordUserRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação: "+err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now()

	updatePasswordParams := db.UpdatePasswordParams{
		ID:        updatePasswordUserRequest.ID,
		Password:  updatePasswordUserRequest.Password,
		UpdatedAt: now,
	}

	if err := queries.UpdatePassword(r.Context(), updatePasswordParams); err != nil {
		http.Error(w, "Erro ao alterar a senha do usuário", http.StatusInternalServerError)
		return
	}

	// Encode user data in JSON format and send it as a response
	if err := json.NewEncoder(w).Encode(updatePasswordParams); err != nil {
		http.Error(w, "Erro ao codificar a resposta", http.StatusBadRequest)
		return
	}

}

func DisableUserController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var disableUserRequest UserRequest

	// Decodes the request body to obtain the user ID
	if err := json.NewDecoder(r.Body).Decode(&disableUserRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Validates that user ID was provided
	if disableUserRequest.ID == uuid.Nil {
		http.Error(w, "Inválido o ID do usuário!", http.StatusBadRequest)
		return
	}

	// Updates the user's IsActive field to false
	updateUserParams := db.UpdateUserParams{
		ID:        disableUserRequest.ID,
		IsActive:  false,
		UpdatedAt: time.Now(),
	}

	// Performs the update operation on the database
	if err := queries.UpdateUser(r.Context(), updateUserParams); err != nil {
		http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Usuário desabilitado com sucesso",
	}

	// Serialize the response to JSON and write it to the HTTP response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao codificar resposta: %v", err), http.StatusInternalServerError)
		return
	}

}
