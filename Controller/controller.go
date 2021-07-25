package Controller

import (
	"encoding/json"
	"net/http"

	"github.com/JuanJTorres11/restaurant-api/Model"
	"github.com/go-chi/chi/v5"
)

func LoadData(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")
	if date != "" {
		w.Write([]byte("loading..." + date))
	} else {
		w.Write([]byte("loading..."))
	}
}

func ListBuyers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp, err := Model.GetBuyers()
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
