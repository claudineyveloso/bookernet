package attendance

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateAttendance(attendance types.AttendancePayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	attendance.ID = uuid.New()
	now := time.Now()
	attendance.CreatedAt = now
	attendance.UpdatedAt = now

	createAttendanceParams := db.CreateAttendanceParams{
		ID:            attendance.ID,
		DateService:   attendance.DateService,
		StartService:  attendance.StartService,
		EndService:    attendance.EndService,
		Status:        attendance.Status,
		Reminder:      attendance.Reminder,
		OwnerID:       attendance.OwnerID,
		TypeServiceID: attendance.TypeServiceID,
		CreatedAt:     attendance.CreatedAt,
		UpdatedAt:     attendance.UpdatedAt,
	}

	if err := queries.CreateAttendance(ctx, createAttendanceParams); err != nil {
		//http.Error(_, "Erro ao criar usuário", http.StatusInternalServerError)
		fmt.Println("Erro ao criar Bucket:", err)
		return err
	}
	return nil
}

func (s *Store) GetAttendances() ([]*types.Attendance, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbAttendances, err := queries.GetAttendances(ctx)
	if err != nil {
		return nil, err
	}

	var attendances []*types.Attendance
	for _, dbBucket := range dbAttendances {
		attendance := convertDBAttendanceToAttendance(dbBucket)
		attendances = append(attendances, attendance)
	}
	return attendances, nil
}

func (s *Store) GetAttendanceByID(attendancesID uuid.UUID) (*types.Attendance, error) {
	queries := db.New(s.db)
	ctx := context.Background()
	dbAttendance, err := queries.GetAttendance(ctx, attendancesID)
	if err != nil {
		return nil, err
	}
	attendance := convertDBAttendanceToAttendance(dbAttendance)

	return attendance, nil

}

func convertDBAttendanceToAttendance(dbAttendance db.Attendance) *types.Attendance {
	attendance := &types.Attendance{
		ID:            dbAttendance.ID,
		DateService:   dbAttendance.DateService,
		StartService:  dbAttendance.StartService,
		EndService:    dbAttendance.EndService,
		Status:        dbAttendance.Status,
		Reminder:      dbAttendance.Reminder,
		OwnerID:       dbAttendance.OwnerID,
		TypeServiceID: dbAttendance.TypeServiceID,
		CreatedAt:     dbAttendance.CreatedAt,
		UpdatedAt:     dbAttendance.UpdatedAt,
	}
	return attendance
}
