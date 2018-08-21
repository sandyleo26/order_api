package order

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sandyleo26/lalamove/database"
)

type DBAdaptor interface {
	Create(*Order) (*Order, error)
	Update(id uint, status Status) (*Order, int, error)
	Get(*GetOptions) ([]*Order, error)
}

type PgAdaptor struct{}

func (*PgAdaptor) Create(order *Order) (*Order, error) {
	db := database.OpenDB()
	result := db.Create(order)
	return order, result.Error
}

func (*PgAdaptor) Update(id uint, status Status) (*Order, int, error) {
	db := database.OpenDB()
	db.Exec("set transaction isolation level serializable")
	tx := db.Begin()
	var theOrder Order
	if err := tx.Raw("SELECT * FROM order_record WHERE id = ? for update", id).Scan(&theOrder).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	if theOrder.Status == StatusUNASSIGN {
		theOrder.Status = StatusSUCCESS
	} else if theOrder.Status == StatusSUCCESS {
		log.Printf("order %d is already taken\n", id)
		tx.Rollback()
		return nil, http.StatusConflict, fmt.Errorf("ORDER_ALREADY_BEEN_TAKEN")
	}

	if err := tx.Save(&theOrder).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}
	tx.Commit()
	return &theOrder, http.StatusOK, nil
}

func (*PgAdaptor) Get(options *GetOptions) ([]*Order, error) {
	db := database.OpenDB()

	var orders []*Order
	result := db.Offset((options.Page - 1) * options.Limit).Limit(options.Limit).Find(&orders)
	return orders, result.Error
}
