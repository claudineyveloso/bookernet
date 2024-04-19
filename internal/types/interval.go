package types

import (
	"time"

	"github.com/google/uuid"
)

type Interval struct {
	ID              uuid.UUID `json:"id"`
	OwnerID         uuid.UUID `json:"owner_id"`
	IntervalMinutes int32     `json:"interval_minutes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type IntervalPayload struct {
	ID              uuid.UUID `json:"id"`
	OwnerID         uuid.UUID `json:"owner_id" validate:"required"`
	IntervalMinutes int32     `json:"interval_minutes" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type IntervalStore interface {
	CreateInterval(IntervalPayload) error
	GetIntervals() ([]*Interval, error)
	GetIntervalByID(id uuid.UUID) (*Interval, error)
}
