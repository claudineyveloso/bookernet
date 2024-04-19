package types

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID            uuid.UUID `json:"id"`
	DateService   time.Time `json:"date_service"`
	StartService  time.Time `json:"start_service"`
	EndService    time.Time `json:"end_service"`
	Status        string    `json:"status"`
	Reminder      int32     `json:"reminder"`
	OwnerID       uuid.UUID `json:"owner_id"`
	TypeServiceID uuid.UUID `json:"type_service_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AttendancePayload struct {
	ID            uuid.UUID `json:"id"`
	DateService   time.Time `json:"date_service" validate:"required"`
	StartService  time.Time `json:"start_service" validate:"required"`
	EndService    time.Time `json:"end_service" validate:"required"`
	Status        string    `json:"status" validate:"required"`
	Reminder      int32     `json:"reminder" validate:"required"`
	OwnerID       uuid.UUID `json:"owner_id" validate:"required"`
	TypeServiceID uuid.UUID `json:"type_service_id" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AttendanceStore interface {
	CreateAttendance(AttendancePayload) error
	GetAttendances() ([]*Attendance, error)
	GetAttendanceByID(id uuid.UUID) (*Attendance, error)
}
