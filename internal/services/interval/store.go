package interval

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

func (s *Store) CreateInterval(interval types.CreateIntervalPayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	interval.ID = uuid.New()
	now := time.Now()
	interval.CreatedAt = now
	interval.UpdatedAt = now

	createIntervalParams := db.CreateIntervalParams{
		ID:              interval.ID,
		OwnerID:         interval.OwnerID,
		IntervalMinutes: interval.IntervalMinutes,
		CreatedAt:       interval.CreatedAt,
		UpdatedAt:       interval.UpdatedAt,
	}

	if err := queries.CreateInterval(ctx, createIntervalParams); err != nil {
		fmt.Println("Erro ao criar o Intervalo:", err)
		return err
	}
	return nil
}

func (s *Store) GetIntervals() ([]*types.Interval, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbIntervals, err := queries.GetIntervals(ctx)
	if err != nil {
		return nil, err
	}

	var intervals []*types.Interval
	for _, dbInterval := range dbIntervals {
		interval := convertDBIntervalToInterval(dbInterval)
		intervals = append(intervals, interval)
	}
	return intervals, nil
}

func (s *Store) GetIntervalByID(intervalID uuid.UUID) (*types.Interval, error) {
	queries := db.New(s.db)
	ctx := context.Background()
	dbInterval, err := queries.GetInterval(ctx, intervalID)
	if err != nil {
		return nil, err
	}
	user := convertDBIntervalToInterval(dbInterval)

	return user, nil

}

func convertDBIntervalToInterval(dbInterval db.Interval) *types.Interval {
	interval := &types.Interval{
		ID:              dbInterval.ID,
		OwnerID:         dbInterval.OwnerID,
		IntervalMinutes: dbInterval.IntervalMinutes,
		CreatedAt:       dbInterval.CreatedAt,
		UpdatedAt:       dbInterval.UpdatedAt,
	}
	return interval
}
