package insurance

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

func (s *Store) CreateInsurance(insurance types.CreateInsurancePayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	insurance.ID = uuid.New()
	now := time.Now()
	insurance.CreatedAt = now
	insurance.UpdatedAt = now

	createInsuranceParams := db.CreateInsuranceParams{
		ID:        insurance.ID,
		Name:      insurance.Name,
		Period:    insurance.Period,
		CreatedAt: insurance.CreatedAt,
		UpdatedAt: insurance.UpdatedAt,
	}

	if err := queries.CreateInsurance(ctx, createInsuranceParams); err != nil {
		//http.Error(_, "Erro ao criar usuário", http.StatusInternalServerError)
		fmt.Println("Erro ao criar convênios:", err)
		return err
	}
	return nil
}

func (s *Store) GetInsurances() ([]*types.Insurance, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbInsurances, err := queries.GetInsurances(ctx)
	if err != nil {
		return nil, err
	}

	var insurances []*types.Insurance
	for _, dbInsurance := range dbInsurances {
		insurance := convertDBInsuranceToInsurance(dbInsurance)
		insurances = append(insurances, insurance)
	}
	return insurances, nil
}

func (s *Store) GetInsuranceByID(insuranceID uuid.UUID) (*types.Insurance, error) {
	queries := db.New(s.db)
	ctx := context.Background()
	dbInsurance, err := queries.GetInsurance(ctx, insuranceID)
	if err != nil {
		return nil, err
	}
	insurance := convertDBInsuranceToInsurance(dbInsurance)

	return insurance, nil

}

func convertDBInsuranceToInsurance(dbInsurance db.Insurance) *types.Insurance {
	insurance := &types.Insurance{
		ID:        dbInsurance.ID,
		Name:      dbInsurance.Name,
		Period:    dbInsurance.Period,
		CreatedAt: dbInsurance.CreatedAt,
		UpdatedAt: dbInsurance.UpdatedAt,
	}
	return insurance
}
