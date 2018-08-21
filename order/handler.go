package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewCreateOrderHandler(uc UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		resp, status, err := uc.CreateOrder(&request)
		if err != nil {
			errResp := ErrorResponse{
				Error: err.Error(),
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func NewTakeOrderHandler(uc UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request TakeRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		resp, status, err := uc.TakeOrder(uint(id), &request)
		if err != nil {
			errResp := ErrorResponse{
				Error: err.Error(),
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func NewGetOrderHandler(uc UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		page, err := strconv.ParseInt(r.Form.Get("page"), 10, 64)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		limit, err := strconv.ParseInt(r.Form.Get("limit"), 10, 64)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		resp, status, err := uc.GetOrder(&GetOptions{
			Page:  int(page),
			Limit: int(limit),
		})
		if err != nil {
			errResp := ErrorResponse{
				Error: err.Error(),
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
