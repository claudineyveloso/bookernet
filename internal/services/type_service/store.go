package typeservice

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

func (s *Store) CreateTypeService(type_service types.TypeServicePayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	type_service.ID = uuid.New()
	now := time.Now()
	type_service.CreatedAt = now
	type_service.UpdatedAt = now

	createTypeServiceParams := db.CreateTypeServiceParams{
		ID:        type_service.ID,
		Name:      type_service.Name,
		Duration:  type_service.Duration,
		CreatedAt: type_service.CreatedAt,
		UpdatedAt: type_service.UpdatedAt,
	}

	if err := queries.CreateTypeService(ctx, createTypeServiceParams); err != nil {
		//http.Error(_, "Erro ao criar usuário", http.StatusInternalServerError)
		fmt.Println("Erro ao criar tipo de serviço:", err)
		return err
	}
	return nil
}

func (s *Store) GetTypeServices() ([]*types.TypeService, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbTypeServices, err := queries.GetTypeServices(ctx)
	if err != nil {
		return nil, err
	}

	var typeServices []*types.TypeService
	for _, dbTypeService := range dbTypeServices {
		typeService := convertDBTypeServiceToTypeService(dbTypeService)
		typeServices = append(typeServices, typeService)
	}
	return typeServices, nil
}

func (s *Store) GetTypeServiceByID(typeServicesID uuid.UUID) (*types.TypeService, error) {
	queries := db.New(s.db)
	ctx := context.Background()
	dbTypeService, err := queries.GetTypeService(ctx, typeServicesID)
	if err != nil {
		return nil, err
	}
	type_service := convertDBTypeServiceToTypeService(dbTypeService)

	return type_service, nil

}

func convertDBTypeServiceToTypeService(dbTypeService db.TypeService) *types.TypeService {
	typeService := &types.TypeService{
		ID:        dbTypeService.ID,
		Name:      dbTypeService.Name,
		Duration:  dbTypeService.Duration,
		CreatedAt: dbTypeService.CreatedAt,
		UpdatedAt: dbTypeService.UpdatedAt,
	}
	return typeService
}
