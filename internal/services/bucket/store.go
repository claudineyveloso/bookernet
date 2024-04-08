package bucket

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

func (s *Store) CreateBucket(bucket types.CreateBucketPayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	bucket.ID = uuid.New()
	now := time.Now()
	bucket.CreatedAt = now
	bucket.UpdatedAt = now

	createBucketParams := db.CreateBucketParams{
		ID:                 bucket.ID,
		Description:        bucket.Description,
		Name:               bucket.Name,
		AwsAccessKeyID:     bucket.AwsAccessKeyID,
		AwsSecretAccessKey: bucket.AwsSecretAccessKey,
		AwsRegion:          bucket.AwsRegion,
		CreatedAt:          bucket.CreatedAt,
		UpdatedAt:          bucket.UpdatedAt,
	}

	if err := queries.CreateBucket(ctx, createBucketParams); err != nil {
		//http.Error(_, "Erro ao criar usuário", http.StatusInternalServerError)
		fmt.Println("Erro ao criar Bucket:", err)
		return err
	}
	return nil
}

func (s *Store) GetBuckets() ([]*types.Bucket, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbBuckets, err := queries.GetBuckets(ctx)
	if err != nil {
		return nil, err
	}

	var buckets []*types.Bucket
	for _, dbBucket := range dbBuckets {
		bucket := convertDBBucketToBucket(dbBucket)
		buckets = append(buckets, bucket)
	}
	return buckets, nil
}

func (s *Store) GetBucketByID(bucketID uuid.UUID) (*types.Bucket, error) {
	queries := db.New(s.db)
	ctx := context.Background()
	dbBucket, err := queries.GetBucket(ctx, bucketID)
	if err != nil {
		return nil, err
	}
	bucket := convertDBBucketToBucket(dbBucket)

	return bucket, nil

}

// func scanRowsIntoUser(rows *sql.Rows) (*types.Bucket, error) {
// 	bucket := new(types.Bucket)

// 	err := rows.Scan(
// 		&bucket.ID,
// 		&bucket.Description,
// 		&bucket.Name,
// 		&bucket.AwsAccessKeyID,
// 		&bucket.AwsSecretAccessKey,
// 		&bucket.AwsRegion,
// 		&bucket.CreatedAt,
// 		&bucket.UpdatedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return bucket, nil
// }

func convertDBBucketToBucket(dbBucket db.Bucket) *types.Bucket {
	bucket := &types.Bucket{
		ID:                 dbBucket.ID,
		Description:        dbBucket.Description,
		Name:               dbBucket.Name,
		AwsAccessKeyID:     dbBucket.AwsAccessKeyID,
		AwsSecretAccessKey: dbBucket.AwsSecretAccessKey,
		AwsRegion:          dbBucket.AwsRegion,
		CreatedAt:          dbBucket.CreatedAt,
		UpdatedAt:          dbBucket.UpdatedAt,
	}
	return bucket
}
