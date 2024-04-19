package types

import (
	"time"

	"github.com/google/uuid"
)

type TypeService struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Duration  int32     `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TypeServicePayload struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Duration  int32     `json:"duration" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TypeServiceStore interface {
	CreateTypeService(TypeServicePayload) error
	GetTypeServices() ([]*TypeService, error)
	GetTypeServiceByID(id uuid.UUID) (*TypeService, error)
}
