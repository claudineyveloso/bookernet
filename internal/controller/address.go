package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/google/uuid"
)

type AddressRequest struct {
	ID              uuid.UUID `json:"id"`
	PublicPlace     string    `json:"public_place"`
	Complement      string    `json:"complement"`
	Neighborhood    string    `json:"neighborhood"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	ZipCode         string    `json:"zip_code"`
	AddressableID   uuid.UUID `json:"addressable_id"`
	AddressableType string    `json:"addressable_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func CreateAddressController(w http.ResponseWriter, r *http.Request, queries *db.Queries, ownerID uuid.UUID) error {
	var createAddressRequest AddressRequest
	if err := json.NewDecoder(r.Body).Decode(&createAddressRequest); err != nil {
		return fmt.Errorf("Erro ao decodificar corpo da solicitação: %v", err)
	}

	now := time.Now()
	createAddressParams := CreateAddressParamsFromRequest(createAddressRequest, ownerID, now)

	if err := queries.CreateAddress(r.Context(), createAddressParams); err != nil {
		return fmt.Errorf("Erro ao criar o endereço: %v", err)
	}
	return nil
}
func CreateAddressParamsFromRequest(req AddressRequest, ownerID uuid.UUID, now time.Time) db.CreateAddressParams {
	return db.CreateAddressParams{
		ID:              uuid.New(),
		PublicPlace:     utils.CreateNullString(req.PublicPlace),
		Complement:      utils.CreateNullString(req.Complement),
		Neighborhood:    utils.CreateNullString(req.Neighborhood),
		City:            utils.CreateNullString(req.City),
		State:           utils.CreateNullString(req.State),
		ZipCode:         utils.CreateNullString(req.ZipCode),
		AddressableID:   ownerID,
		AddressableType: "owner",
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
