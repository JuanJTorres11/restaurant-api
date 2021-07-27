package Controller

import (
	"encoding/json"
	"net/http"

	"github.com/JuanJTorres11/restaurant-api/Model"
	"github.com/go-chi/chi/v5"
)

type GetBuyerResponse struct {
	Buyer               Model.QueryBuyer `json:"Buyer"`
	OtherBuyers         []string         `json:"Other_Buyers"`
	RecommendedProducts []string         `json:"Recommended_Products"`
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	date := chi.URLParam(r, "date")
	_, err1 := Model.GetBuyers(date)
	if err1 != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{
		"Status": "Error",
		"Message": "There was an error retrieving the data of the buyers"
		}
		`))
		return
	}
	_, err2 := Model.GetProducts(date)
	if err2 != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{
		"Status": "Error",
		"Message": "There was an error retrieving the data of the products"
		}
		`))
		return
	}
	_, err3 := Model.GetTransactions(date)
	if err3 != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{
		"Status": "Error",
		"Message": "There was an error retrieving the data of the transactions"
		}
		`))
		return
	}
	w.Write([]byte(`{
		"Status": "Success",
		"Message": "All the data has been uploaded successfully"
		}
		`))
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := chi.URLParam(r, "buyerID")
	buyer, buyerNames, productNames, err := Model.GetBuyer(id)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{
		"Status": "Error",
		"Message": "There was an error retrieving the data"
		}
		`))
		return
	}

	response := GetBuyerResponse{buyer, buyerNames, productNames}

	json.NewEncoder(w).Encode(response)
}
