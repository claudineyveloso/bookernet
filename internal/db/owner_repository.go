package db

import (
	"context"

	"github.com/google/uuid"
)

// Defina uma nova struct para representar os dados completos do cliente
type FullOwnerRow struct {
	Owner   GetOwnerRow
	Person  Person
	Address Address
}

func (q *Queries) GetOwnersWithDetails(ctx context.Context) ([]FullOwnerRow, error) {
	rows, err := q.db.QueryContext(ctx, getOwners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var owners []FullOwnerRow
	for rows.Next() {
		var owner FullOwnerRow
		err := rows.Scan(
			&owner.Owner.ID,
			&owner.Owner.PeopleType,
			&owner.Owner.IsActive,
			&owner.Owner.BucketID,
			&owner.Owner.CreatedAt,
			&owner.Owner.UpdatedAt,
			&owner.Person.ID,
			&owner.Person.FirstName,
			&owner.Person.LastName,
			&owner.Person.Email,
			&owner.Person.Phone,
			&owner.Person.CellPhone,
			&owner.Person.PersonableID,
			&owner.Person.PersonableType,
			&owner.Person.CreatedAt,
			&owner.Person.UpdatedAt,
			&owner.Address.ID,
			&owner.Address.PublicPlace,
			&owner.Address.Complement,
			&owner.Address.Neighborhood,
			&owner.Address.City,
			&owner.Address.State,
			&owner.Address.ZipCode,
			&owner.Address.AddressableID,
			&owner.Address.AddressableType,
			&owner.Address.CreatedAt,
			&owner.Address.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return owners, nil
}

func (q *Queries) GetOwnerWithDetails(ctx context.Context, id uuid.UUID) (FullOwnerRow, error) {
	var owner FullOwnerRow
	err := q.db.QueryRowContext(ctx, getOwner, id).Scan(
		&owner.Owner.ID,
		&owner.Owner.PeopleType,
		&owner.Owner.IsActive,
		&owner.Owner.BucketID,
		&owner.Owner.CreatedAt,
		&owner.Owner.UpdatedAt,
		&owner.Person.ID,
		&owner.Person.FirstName,
		&owner.Person.LastName,
		&owner.Person.Email,
		&owner.Person.Phone,
		&owner.Person.CellPhone,
		&owner.Person.PersonableID,
		&owner.Person.PersonableType,
		&owner.Person.CreatedAt,
		&owner.Person.UpdatedAt,
		&owner.Address.ID,
		&owner.Address.PublicPlace,
		&owner.Address.Complement,
		&owner.Address.Neighborhood,
		&owner.Address.City,
		&owner.Address.State,
		&owner.Address.ZipCode,
		&owner.Address.AddressableID,
		&owner.Address.AddressableType,
		&owner.Address.CreatedAt,
		&owner.Address.UpdatedAt,
	)
	if err != nil {
		return FullOwnerRow{}, err
	}
	return owner, nil
}
