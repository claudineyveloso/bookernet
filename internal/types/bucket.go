package types

import (
	"time"

	"github.com/google/uuid"
)

type Bucket struct {
	ID                 uuid.UUID `json:"id"`
	Description        string    `json:"description"`
	Name               string    `json:"name"`
	AwsAccessKeyID     string    `json:"aws_access_key_id"`
	AwsSecretAccessKey string    `json:"aws_secret_access_key"`
	AwsRegion          string    `json:"aws_region"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type BucketPayload struct {
	ID                 uuid.UUID `json:"id"`
	Description        string    `json:"description"`
	Name               string    `json:"name" validate:"required"`
	AwsAccessKeyID     string    `json:"aws_access_key_id" validate:"required"`
	AwsSecretAccessKey string    `json:"aws_secret_access_key" validate:"required"`
	AwsRegion          string    `json:"aws_region" validate:"required"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type BucketStore interface {
	CreateBucket(BucketPayload) error
	GetBuckets() ([]*Bucket, error)
	GetBucketByID(id uuid.UUID) (*Bucket, error)
	UpdateBucket(BucketPayload) error
	DeleteBucket(id uuid.UUID) error
}
