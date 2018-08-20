package order

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestCreateOrderHandler(t *testing.T) {

	mockUC := MockUseCase{}

	handler := NewCreateOrderHandler(&mockUC)

	cr := CreateRequest{
		Origin:      Location{"11", "22"},
		Destination: Location{"33", "44"},
	}

	marshaled, _ := json.Marshal(cr)
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(marshaled))
	rr := httptest.NewRecorder()

	mockUC.On("CreateOrder", mock.Anything).Return(&CreateResponse{}, http.StatusOK, nil)
	handler(rr, req)
	mockUC.AssertCalled(t, "CreateOrder")
	assert.Equal(t, http.StatusOK, rr.Code)
}
