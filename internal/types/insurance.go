package types

import (
	"time"

	"github.com/google/uuid"
)

type Insurance struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Period    string    `json:"period"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsurancePayload struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Period    string    `json:"period" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsuranceStore interface {
	CreateInsurance(InsurancePayload) error
	GetInsurances() ([]*Insurance, error)
	GetInsuranceByID(id uuid.UUID) (*Insurance, error)
}
