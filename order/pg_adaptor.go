package order

import "github.com/sandyleo26/lalamove/database"

type OutputAdaptor interface {
	Create(*Order) (*Order, error)
	Update(*Order) (*Order, error)
}

type PgAdaptor struct{}

func (*PgAdaptor) Create(order *Order) (*Order, error) {
	db := database.OpenDB()
	result := db.Create(order)
	return order, result.Error
}

func (*PgAdaptor) Update(order *Order) (*Order, error) {
	return nil, nil
}
