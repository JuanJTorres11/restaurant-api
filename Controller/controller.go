package Controller

import (
	"encoding/json"
	"net/http"

	"github.com/JuanJTorres11/restaurant-api/Model"
	"github.com/go-chi/chi/v5"
)

func LoadData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	date := chi.URLParam(r, "date")
	// resp, err := Model.GetBuyers(date)
	// resp, err := Model.GetProducts(date)
	resp, err := Model.GetTransactions(date)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{
		"Status": "Error",
		"Message": "There was an error retrieving the data"
		}
		`))
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func ListBuyers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp, err := Model.ListBuyers()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{
		"Status": "Error",
		"Message": "There was an error retrieving the data"
		}
		`))
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func GetBuyer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "buyerID")
	w.Write([]byte("loading..." + id))
}
