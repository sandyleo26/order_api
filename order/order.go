package order

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Location []string

//Order is the main model
type CreateRequest struct {
	Origin      Location
	Destination Location
}

func (r CreateRequest) validate() error {
	if len(r.Origin) < 2 {
		return fmt.Errorf("Invalid origin")
	}

	if len(r.Destination) < 2 {
		return fmt.Errorf("Invalid destination")
	}
	return nil
}

type Status string

const (
	StatusUNASSIGN Status = "UNASSIGN"
	StatusSUCCESS  Status = "SUCCESS"
)

type CreateResponse struct {
	ID       uint
	Distance int
	Status
}

type TakeRequest struct {
	Status
}

func (r TakeRequest) validate() error {
	if r.Status != "taken" {
		return fmt.Errorf("Invalid TakeRequest")
	}
	return nil
}

type TakeResponse struct {
	Status
}

type GetOptions struct {
	Page  int
	Limit int
}

func (opt GetOptions) validate() error {
	if opt.Limit < 0 || opt.Page < 0 {
		return fmt.Errorf("invalid GetOptions")
	}
	return nil
}

type GetResponse CreateResponse

type ErrorResponse struct {
	Error string
}

//Order gorm model
type Order struct {
	gorm.Model
	StartLat  string
	StartLong string
	EndLat    string
	EndLong   string
	Distance  int
	Status    Status
}

func (*Order) TableName() string {
	return "order_record"
}
