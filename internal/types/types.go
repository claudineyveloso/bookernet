package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Address struct {
	ID              uuid.UUID `json:"id"`
	PublicPlace     string    `json:"public_place"`
	Complement      string    `json:"complement"`
	Neighborhood    string    `json:"neighborhood"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	ZipCode         string    `json:"zip_code"`
	AddressableID   uuid.UUID `json:"addressable_id"`
	AddressableType string    `json:"addressable_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Person struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	CellPhone      string    `json:"cell_phone"`
	PersonableID   uuid.UUID `json:"personable_id"`
	PersonableType string    `json:"personable_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Owner struct {
	ID         uuid.UUID `json:"id"`
	PeopleType string    `json:"people_type"`
	IsActive   bool      `json:"is_active"`
	BucketID   uuid.UUID `json:"bucket_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Person     `json:"person"`
	Address    `json:"address"`
}

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

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Person    `json:"person"`
	Address   `json:"address"`
}

type CreateUserPayload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	IsActive  bool      `json:"is_active"`
	UserType  string    `json:"user_type" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAddressPayload struct {
	ID              uuid.UUID `json:"id"`
	PublicPlace     string    `json:"public_place"`
	Complement      string    `json:"complement"`
	Neighborhood    string    `json:"neighborhood"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	ZipCode         string    `json:"zip_code"`
	AddressableID   uuid.UUID `json:"addressable_id"`
	AddressableType string    `json:"addressable_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreatePersonPayload struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	CellPhone      string    `json:"cell_phone" validate:"required"`
	PersonableID   uuid.UUID `json:"personable_id"`
	PersonableType string    `json:"personable_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateOwnerPayload struct {
	ID         uuid.UUID `json:"id"`
	PeopleType string    `json:"people_type" validate:"required"`
	IsActive   bool      `json:"is_active"`
	BucketID   uuid.UUID `json:"bucket_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Person     `json:"person"`
	Address    `json:"address"`
}
type CreateBucketPayload struct {
	ID                 uuid.UUID `json:"id"`
	Description        string    `json:"description" validate:"required"`
	Name               string    `json:"name" validate:"required"`
	AwsAccessKeyID     string    `json:"aws_access_key_id" validate:"required"`
	AwsSecretAccessKey string    `json:"aws_secret_access_key" validate:"required"`
	AwsRegion          string    `json:"aws_region" validate:"required"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateCustomerPayload struct {
	ID        uuid.UUID `json:"id"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Person    `json:"person"`
	Address   `json:"address"`
}

type PasswordUserPayload struct {
	Password  string    `json:"password" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserPayload struct {
	Email     string    `json:"email" validate:"required"`
	IsActive  bool      `json:"is_active"`
	UserType  string    `json:"user_type" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserStore interface {
	CreateUser(CreateUserPayload) error
	GetUsers() ([]*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	//UpdateUser(User) error
	//UpdateUser(User) error
}

type AddressStore interface {
	CreateAddress(CreateAddressPayload) error
}

type PersonStore interface {
	CreatePerson(CreatePersonPayload) error
}

type OwnerStore interface {
	CreateOwner(CreateOwnerPayload) (uuid.UUID, error)
}

type BucketStore interface {
	CreateBucket(CreateBucketPayload) error
	GetBuckets() ([]*Bucket, error)
	GetBucketByID(id uuid.UUID) (*Bucket, error)
}

type CustomerStore interface {
	CreateCustomer(CreateCustomerPayload) (uuid.UUID, error)
}
