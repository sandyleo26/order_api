package order

import (
	"fmt"
	"log"
	"net/http"
)

type UseCase interface {
	CreateOrder(r *CreateRequest) (*CreateResponse, int, error)
	TakeOrder(id int, r *TakeRequest) (*TakeResponse, int, error)
	GetOrder(options *GetOptions) ([]*GetResponse, int, error)
}

type RealUseCase struct {
	OutputAdaptor
	DistanceService
}

func NewRealUseCase() UseCase {
	return &RealUseCase{
		OutputAdaptor:   &PgAdaptor{},
		DistanceService: &GoogleDistanceService{},
	}
}

func (uc *RealUseCase) CreateOrder(r *CreateRequest) (*CreateResponse, int, error) {
	if err := r.validate(); err != nil {
		log.Printf("validate failed with error=%s\n", err.Error())
		return nil, http.StatusInternalServerError, fmt.Errorf("ERROR_DESCRIPTION")
	}

	d, err := uc.DistanceService.Calculate(r.Origin, r.Destination)
	if err != nil {
		log.Printf("Calculate failed with error=%s\n", err.Error())
		return nil, http.StatusInternalServerError, fmt.Errorf("ERROR_DESCRIPTION")
	}

	newOrder := &Order{
		StartLat:  r.Origin[0],
		StartLong: r.Origin[1],
		EndLat:    r.Destination[0],
		EndLong:   r.Destination[1],
		Distance:  d,
		Status:    StatusUNASSIGN,
	}
	_, err = uc.OutputAdaptor.Create(newOrder)
	if err != nil {
		log.Printf("Create order failed with error=%s\n", err.Error())
		return nil, http.StatusInternalServerError, fmt.Errorf("ERROR_DESCRIPTION")
	}

	return &CreateResponse{
		ID:       newOrder.ID,
		Distance: newOrder.Distance,
		Status:   newOrder.Status,
	}, http.StatusOK, nil
}

func (*RealUseCase) TakeOrder(id int, r *TakeRequest) (*TakeResponse, int, error) {
	return &TakeResponse{
		Status: "SUCCESS",
	}, http.StatusOK, nil
}

func (*RealUseCase) GetOrder(options *GetOptions) ([]*GetResponse, int, error) {
	return []*GetResponse{}, http.StatusOK, nil
}
