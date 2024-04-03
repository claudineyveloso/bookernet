package db

import "context"

// Defina uma nova struct para representar os dados completos do cliente
type FullCustomerRow struct {
	Customer GetCustomerRow
	Person   Person
	Address  Address
}

func (q *Queries) GetCustomersWithDetails(ctx context.Context) ([]FullCustomerRow, error) {
	rows, err := q.db.QueryContext(ctx, getCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []FullCustomerRow
	for rows.Next() {
		var customer FullCustomerRow
		err := rows.Scan(
			&customer.Customer.ID,
			&customer.Customer.Birthday,
			&customer.Customer.CreatedAt,
			&customer.Customer.UpdatedAt,
			&customer.Person.ID,
			&customer.Person.FirstName,
			&customer.Person.LastName,
			&customer.Person.Email,
			&customer.Person.Phone,
			&customer.Person.CellPhone,
			&customer.Person.PersonableID,
			&customer.Person.PersonableType,
			&customer.Person.CreatedAt,
			&customer.Person.UpdatedAt,
			&customer.Address.ID,
			&customer.Address.PublicPlace,
			&customer.Address.Complement,
			&customer.Address.Neighborhood,
			&customer.Address.City,
			&customer.Address.State,
			&customer.Address.ZipCode,
			&customer.Address.AddressableID,
			&customer.Address.AddressableType,
			&customer.Address.CreatedAt,
			&customer.Address.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}
