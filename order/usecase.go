package order

import (
	"fmt"
	"log"
	"net/http"
)

type UseCase interface {
	CreateOrder(r *CreateRequest) (*CreateResponse, int, error)
	TakeOrder(id uint, r *TakeRequest) (*TakeResponse, int, error)
	GetOrder(options *GetOptions) ([]*GetResponse, int, error)
}

type RealUseCase struct {
	DBAdaptor
	DistanceService
}

func NewRealUseCase() UseCase {
	return &RealUseCase{
		DBAdaptor:       &PgAdaptor{},
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
	newOrder, err = uc.DBAdaptor.Create(newOrder)
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

func (uc *RealUseCase) TakeOrder(id uint, r *TakeRequest) (*TakeResponse, int, error) {
	if err := r.validate(); err != nil {
		log.Printf("TakeRequest validate failed with error=%s", err.Error())
		return nil, http.StatusBadRequest, fmt.Errorf("bad request")
	}

	_, err := uc.DBAdaptor.Update(id, StatusSUCCESS)
	if err != nil {

	}
	return &TakeResponse{
		Status: "SUCCESS",
	}, http.StatusOK, nil
}

func (uc *RealUseCase) GetOrder(options *GetOptions) ([]*GetResponse, int, error) {
	if err := options.validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	orders, err := uc.DBAdaptor.Get(options)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	responses := make([]*GetResponse, 0)
	for _, order := range orders {
		resp := &GetResponse{
			ID:       order.ID,
			Distance: order.Distance,
			Status:   order.Status,
		}
		responses = append(responses, resp)
	}

	return responses, http.StatusOK, nil
}
