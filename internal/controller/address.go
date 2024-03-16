package controller

import (
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

func CreateAddressController(w http.ResponseWriter, r *http.Request, queries *db.Queries, createAddressRequest *OwnerRequest) (uuid.UUID, error) {
	now := time.Now()

	createAddressParams := db.CreateAddressParams{
		ID:              uuid.New(),
		PublicPlace:     utils.CreateNullString(createAddressRequest.Address.PublicPlace),
		Complement:      utils.CreateNullString(createAddressRequest.Address.Complement),
		Neighborhood:    utils.CreateNullString(createAddressRequest.Address.Neighborhood),
		City:            utils.CreateNullString(createAddressRequest.Address.City),
		State:           utils.CreateNullString(createAddressRequest.Address.State),
		ZipCode:         utils.CreateNullString(createAddressRequest.Address.ZipCode),
		AddressableID:   createAddressRequest.ID,
		AddressableType: "owner",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	err := queries.CreateAddress(r.Context(), createAddressParams)
	if err != nil {
		return uuid.Nil, err
	}

	return createAddressParams.ID, nil
}
