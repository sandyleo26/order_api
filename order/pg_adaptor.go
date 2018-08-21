package order

import (
	"fmt"

	"github.com/sandyleo26/lalamove/database"
)

type OutputAdaptor interface {
	Create(*Order) (*Order, error)
	Update(id uint, status Status) (*Order, error)
	Get(*GetOptions) ([]*Order, error)
}

type PgAdaptor struct{}

func (*PgAdaptor) Create(order *Order) (*Order, error) {
	db := database.OpenDB()
	result := db.Create(order)
	return order, result.Error
}

func (*PgAdaptor) Update(id uint, status Status) (*Order, error) {
	db := database.OpenDB()
	db.Exec("set transaction isolation level serializable")
	tx := db.Begin()
	var theOrder Order
	if err := tx.Raw("SELECT * FROM order WHERE id = ? for update", id).Scan(&theOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if theOrder.Status == StatusUNASSIGN {
		theOrder.Status = StatusSUCCESS
	} else if theOrder.Status == StatusSUCCESS {
		return nil, fmt.Errorf("order %d is already taken", id)
	}

	if err := tx.Save(&theOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &theOrder, nil
}

func (*PgAdaptor) Get(options *GetOptions) ([]*Order, error) {
	db := database.OpenDB()

	var orders []*Order
	result := db.Offset((options.Page - 1) * options.Limit).Limit(options.Limit).Find(&orders)
	return orders, result.Error
}
