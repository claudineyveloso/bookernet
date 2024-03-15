package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BucketRequest struct {
	ID                 uuid.UUID `json:"id"`
	Description        string    `json:"description" validate:"required"`
	Name               string    `json:"name" validate:"required"`
	AwsAccessKeyID     string    `json:"aws_access_key_id" validate:"required"`
	AwsSecretAccessKey string    `json:"aws_secret_access_key" validate:"required"`
	AwsRegion          string    `json:"aws_region" validate:"required"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func CreateBucketController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var createBucketRequest BucketRequest

	if err := json.NewDecoder(r.Body).Decode(&createBucketRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitaçãp", http.StatusBadRequest)
		return
	}

	createBucketRequest.ID = uuid.New()
	now := time.Now()
	createBucketRequest.CreatedAt = now
	createBucketRequest.UpdatedAt = now

	createBucketParams := db.CreateBucketParams{
		ID:                 createBucketRequest.ID,
		Description:        createBucketRequest.Description,
		Name:               createBucketRequest.Name,
		AwsAccessKeyID:     createBucketRequest.AwsAccessKeyID,
		AwsSecretAccessKey: createBucketRequest.AwsSecretAccessKey,
		AwsRegion:          createBucketRequest.AwsRegion,
		CreatedAt:          createBucketRequest.CreatedAt,
		UpdatedAt:          createBucketRequest.UpdatedAt,
	}

	if err := queries.CreateBucket(r.Context(), createBucketParams); err != nil {
		http.Error(w, "Erro ao criar bucket", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(createBucketParams); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (u *BucketRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Description":
				errorMap["description"] = "O campo de descri não pode estar vazio"
			case "Name":
				errorMap["name"] = "O campo de nome não pode estar vazio"
			case "AwsAccessKeyID":
				errorMap["aws_access_key_id"] = "O campo de tipo de usuário não pode estar vazio"
			case "AwsSecretAccessKey":
				errorMap["aws_secret_access_key"] = "O campo de tipo de usuário não pode estar vazio"
			case "AwsRegion":
				errorMap["aws_region"] = "O campo de tipo de usuário não pode estar vazio"
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

func GetBucketsController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	buckets, err := queries.GetBuckets(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter bucket: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(buckets, "", " ")
	if err != nil {
		http.Error(w, "Erro ao codificar em JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(
			w,
			"Erro ao escrever resposta",
			http.StatusInternalServerError,
		)
		return
	}

}

func GetBucketController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	bucketID := r.URL.Query().Get("id")
	parsedBucketID, err := uuid.Parse(bucketID)
	if err != nil {
		http.Error(w, "ID do bucket não pode ser vazio", http.StatusBadRequest)
		return
	}
	bucket, err := queries.GetBucket(r.Context(), parsedBucketID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter usuário: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(bucket, "", " ")
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

func UpdateBucketController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var updateBucketRequest BucketRequest
	if err := json.NewDecoder(r.Body).Decode(&updateBucketRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação: "+err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now()

	// Create instance of db.UpdateUserParams based on request data
	updateBucketParams := db.UpdateBucketParams{
		ID:                 updateBucketRequest.ID,
		Description:        updateBucketRequest.Description,
		Name:               updateBucketRequest.Name,
		AwsAccessKeyID:     updateBucketRequest.AwsAccessKeyID,
		AwsSecretAccessKey: updateBucketRequest.AwsSecretAccessKey,
		AwsRegion:          updateBucketRequest.AwsRegion,
		UpdatedAt:          now,
	}

	if err := queries.UpdateBucket(r.Context(), updateBucketParams); err != nil {
		http.Error(w, "Erro ao alterar dados do usuário", http.StatusInternalServerError)
		return
	}

	// Encode user data in JSON format and send it as a response
	if err := json.NewEncoder(w).Encode(updateBucketParams); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}
