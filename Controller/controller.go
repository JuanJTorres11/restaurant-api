package Controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("loading..."))
}

func LoadDataDate(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")
	w.Write([]byte("loading..." + date))
}

func ListBuyers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("loading..."))
}

func GetBuyer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "buyerID")
	w.Write([]byte("loading..." + id))
}
