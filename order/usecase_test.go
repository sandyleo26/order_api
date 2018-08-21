package order

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/gorm"
)

func TestCreateOrder(t *testing.T) {
	mockDBAdaptor := &MockDBAdaptor{}
	mockDistanceService := &MockDistanceService{}
	uc := &RealUseCase{
		DBAdaptor:       mockDBAdaptor,
		DistanceService: mockDistanceService,
	}

	r := &CreateRequest{
		Origin:      Location{"11", "22"},
		Destination: Location{"33", "44"},
	}

	fakeDistance := 10
	mockDistanceService.On("Calculate", r.Origin, r.Destination).Return(fakeDistance, nil)

	fakeOrderID := uint(25)
	mockDBAdaptor.On("Create", &Order{
		StartLat:  r.Origin[0],
		StartLong: r.Origin[1],
		EndLat:    r.Destination[0],
		EndLong:   r.Destination[1],
		Distance:  fakeDistance,
		Status:    StatusUNASSIGN,
	}).Return(&Order{Model: gorm.Model{ID: fakeOrderID}, Distance: fakeDistance}, nil)
	resp, status, err := uc.CreateOrder(r)
	assert.Nil(t, err)
	assert.Equal(t, fakeOrderID, resp.ID)
	assert.Equal(t, fakeDistance, resp.Distance)
	assert.Equal(t, http.StatusOK, status)
	mockDBAdaptor.AssertExpectations(t)
	mockDistanceService.AssertExpectations(t)
}

func TestTakeOrder(t *testing.T) {
	mockDBAdaptor := &MockDBAdaptor{}
	uc := &RealUseCase{
		DBAdaptor: mockDBAdaptor,
	}

	r := &TakeRequest{
		Status: "taken",
	}

	firstOrderID := uint(10)
	mockDBAdaptor.On("Update", firstOrderID, StatusSUCCESS).Return(&Order{
		Model: gorm.Model{
			ID: firstOrderID,
		},
		Status: StatusSUCCESS,
	}, nil)
	resp, status, err := uc.TakeOrder(firstOrderID, r)
	assert.Nil(t, err)
	assert.Equal(t, StatusSUCCESS, resp.Status)
	assert.Equal(t, http.StatusOK, status)
}
