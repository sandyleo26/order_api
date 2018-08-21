package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sandyleo26/order_api/order"
	"github.com/stretchr/testify/mock"
)

func TestRouterSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	b, _ := json.Marshal(order.CreateRequest{
		Origin:      order.Location{"11", "22"},
		Destination: order.Location{"33", "44"},
	})
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(b))
	mockUC := order.MockUseCase{}
	mockUC.On("CreateOrder", mock.Anything).Return(&order.CreateResponse{}, http.StatusOK, nil)
	Router(&mockUC).ServeHTTP(rr, req)
	mockUC.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rr.Code)

	rr = httptest.NewRecorder()
	b, _ = json.Marshal(order.TakeRequest{
		Status: "taken",
	})
	req, _ = http.NewRequest("PUT", "/order/42", bytes.NewBuffer(b))
	mockUC.On("TakeOrder", uint(42), &order.TakeRequest{Status: "taken"}).Return(&order.TakeResponse{}, http.StatusOK, nil)
	Router(&mockUC).ServeHTTP(rr, req)
	mockUC.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rr.Code)

	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/orders?page=1&limit=8", nil)
	mockUC.On("GetOrder", &order.GetOptions{Page: 1, Limit: 8}).Return([]*order.GetResponse{}, http.StatusOK, nil)
	Router(&mockUC).ServeHTTP(rr, req)
	mockUC.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRouterErrorCondition(t *testing.T) {
	// create order wrong method
	mockUC := order.MockUseCase{}
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/order", nil)
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	// wrong route
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/wrong", bytes.NewBuffer([]byte("some bytes")))
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// wrong request body
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/order", bytes.NewBuffer([]byte("some bytes")))
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// internal error
	rr = httptest.NewRecorder()
	cr := &order.CreateRequest{
		Origin:      order.Location{"11", "22"},
		Destination: order.Location{"33", "44"},
	}
	b, _ := json.Marshal(cr)
	req, _ = http.NewRequest("POST", "/order", bytes.NewBuffer(b))
	mockUC.On("CreateOrder", cr).Return(&order.CreateResponse{}, http.StatusInternalServerError, fmt.Errorf("CreateOrder error"))
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	var mockErrorResponse order.ErrorResponse
	json.NewDecoder(rr.Body).Decode(&mockErrorResponse)
	assert.Equal(t, &order.ErrorResponse{Error: "CreateOrder error"}, &mockErrorResponse)

	// take order wrong method
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/order/42", bytes.NewBuffer([]byte("some bytes")))
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	// wrong route
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/orders/42", bytes.NewBuffer([]byte("some bytes")))
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// wrong request body
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/order/42", bytes.NewBuffer([]byte("some bytes")))
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// internal error
	rr = httptest.NewRecorder()
	b, _ = json.Marshal(&order.TakeRequest{})
	req, _ = http.NewRequest("PUT", "/order/42", bytes.NewBuffer(b))
	mockUC.On("TakeOrder", uint(42), &order.TakeRequest{}).Return(&order.TakeResponse{}, http.StatusConflict, fmt.Errorf("TakeOrder error"))
	Router(&mockUC).ServeHTTP(rr, req)
	mockErrorResponse = order.ErrorResponse{}
	json.NewDecoder(rr.Body).Decode(&mockErrorResponse)
	assert.Equal(t, &order.ErrorResponse{Error: "TakeOrder error"}, &mockErrorResponse)
	assert.Equal(t, http.StatusConflict, rr.Code)

	// get order wrong method
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/orders?page=1&limit=8", nil)
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	// wrong route
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/orderssss", nil)
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// wrong query
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/orders?abc", nil)
	Router(&mockUC).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// internal error
	rr = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/orders?page=1&limit=8", nil)
	mockUC.On("GetOrder", &order.GetOptions{Limit: 8, Page: 1}).Return([]*order.GetResponse{}, http.StatusInternalServerError, fmt.Errorf("GetOrder error"))
	Router(&mockUC).ServeHTTP(rr, req)
	mockErrorResponse = order.ErrorResponse{}
	json.NewDecoder(rr.Body).Decode(&mockErrorResponse)
	assert.Equal(t, &order.ErrorResponse{Error: "GetOrder error"}, &mockErrorResponse)
}
