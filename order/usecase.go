package order

import (
	"fmt"
	"net/http"
)

type UseCase interface {
	CreateOrder(r *CreateRequest) (*CreateResponse, int, error)
	TakeOrder(id int, r *TakeRequest) (*TakeResponse, int, error)
	GetOrder(options *GetOptions) ([]*GetResponse, int, error)
}

type RealUseCase struct {
}

func (*RealUseCase) CreateOrder(r *CreateRequest) (*CreateResponse, int, error) {
	fmt.Println(r)
	return &CreateResponse{
		ID:       123,
		Distance: 400,
		Status:   "UNASSIGN",
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
