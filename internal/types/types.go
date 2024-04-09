package types

import (
	"database/sql"
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
	Person     Person    `json:"person"`
	Address    Address   `json:"address"`
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
	Person    Person    `json:"person"`
	Address   Address   `json:"address"`
}

type GetCustomersRow struct {
	ID              uuid.UUID    `json:"id"`
	Birthday        sql.NullTime `json:"birthday"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	ID_2            uuid.UUID    `json:"id_2"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	Email           string       `json:"email"`
	Phone           string       `json:"phone"`
	CellPhone       string       `json:"cell_phone"`
	PersonableID    uuid.UUID    `json:"personable_id"`
	PersonableType  string       `json:"personable_type"`
	CreatedAt_2     time.Time    `json:"created_at_2"`
	UpdatedAt_2     time.Time    `json:"updated_at_2"`
	ID_3            uuid.UUID    `json:"id_3"`
	PublicPlace     string       `json:"public_place"`
	Complement      string       `json:"complement"`
	Neighborhood    string       `json:"neighborhood"`
	City            string       `json:"city"`
	State           string       `json:"state"`
	ZipCode         string       `json:"zip_code"`
	AddressableID   uuid.UUID    `json:"addressable_id"`
	AddressableType string       `json:"addressable_type"`
	CreatedAt_3     time.Time    `json:"created_at_3"`
	UpdatedAt_3     time.Time    `json:"updated_at_3"`
}

type GetOwnerRow struct {
	ID              uuid.UUID `json:"id"`
	PeopleType      string    `json:"people_type"`
	IsActive        bool      `json:"is_active"`
	BucketID        uuid.UUID `json:"bucket_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ID_2            uuid.UUID `json:"id_2"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	CellPhone       string    `json:"cell_phone"`
	PersonableID    uuid.UUID `json:"personable_id"`
	PersonableType  string    `json:"personable_type"`
	CreatedAt_2     time.Time `json:"created_at_2"`
	UpdatedAt_2     time.Time `json:"updated_at_2"`
	ID_3            uuid.UUID `json:"id_3"`
	PublicPlace     string    `json:"public_place"`
	Complement      string    `json:"complement"`
	Neighborhood    string    `json:"neighborhood"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	ZipCode         string    `json:"zip_code"`
	AddressableID   uuid.UUID `json:"addressable_id"`
	AddressableType string    `json:"addressable_type"`
	CreatedAt_3     time.Time `json:"created_at_3"`
	UpdatedAt_3     time.Time `json:"updated_at_3"`
}

type TypeService struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Duration  int32     `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Interval struct {
	ID              uuid.UUID `json:"id"`
	OwnerID         uuid.UUID `json:"owner_id"`
	IntervalMinutes int32     `json:"interval_minutes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

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

type Insurance struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Period    string    `json:"period"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Person     Person    `json:"person"`
	Address    Address   `json:"address"`
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
	Person    Person    `json:"person"`
	Address   Address   `json:"address"`
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

type RegisterUserPayload struct {
	//FirstName string `json:"firstName" validate:"required"`
	//LastName  string `json:"lastName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateTypeServicePayload struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Duration  int32     `json:"duration" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateIntervalPayload struct {
	ID              uuid.UUID `json:"id"`
	OwnerID         uuid.UUID `json:"owner_id" validate:"required"`
	IntervalMinutes int32     `json:"interval_minutes" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateAttendancePayload struct {
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

type CreateInsurancePayload struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Period    string    `json:"period" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserStore interface {
	CreateUser(CreateUserPayload) error
	GetUsers() ([]*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
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
	GetOwners() ([]*Owner, error)
	GetOwner(id uuid.UUID) (*Owner, error)
}

type BucketStore interface {
	CreateBucket(CreateBucketPayload) error
	GetBuckets() ([]*Bucket, error)
	GetBucketByID(id uuid.UUID) (*Bucket, error)
}

type CustomerStore interface {
	CreateCustomer(CreateCustomerPayload) (uuid.UUID, error)
	GetCustomers() ([]*Customer, error)
	GetCustomer(id uuid.UUID) (*Customer, error)
}

type TypeServiceStore interface {
	CreateTypeService(CreateTypeServicePayload) error
	GetTypeServices() ([]*TypeService, error)
	GetTypeServiceByID(id uuid.UUID) (*TypeService, error)
}

type IntervalStore interface {
	CreateInterval(CreateIntervalPayload) error
	GetIntervals() ([]*Interval, error)
	GetIntervalByID(id uuid.UUID) (*Interval, error)
}

type AttendanceStore interface {
	CreateAttendance(CreateAttendancePayload) error
	GetAttendances() ([]*Attendance, error)
	GetAttendanceByID(id uuid.UUID) (*Attendance, error)
}

type InsuranceStore interface {
	CreateInsurance(CreateInsurancePayload) error
	GetInsurances() ([]*Insurance, error)
	GetInsuranceByID(id uuid.UUID) (*Insurance, error)
}
