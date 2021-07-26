package Model

import (
	"bytes"
	"encoding/csv"
	"log"
	"strconv"
)

type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

type ProductID struct {
	ID string `json:"id,omitempty"`
}

func formatProducts(r []byte) []Product {
	var products []Product
	reader := csv.NewReader(bytes.NewReader(r))
	reader.Comma = '\''
	result, err := reader.ReadAll()
	if err != nil {
		log.Panicln("There was an error trying to convert the CSV of the products")
	}

	allKeys := make(map[string]bool)
	for _, item := range result {
		if _, value := allKeys[item[0]]; !value {
			allKeys[item[0]] = true
			p, _ := strconv.Atoi(item[2])
			products = append(products, Product{item[0], item[1], p})
		}
	}

	return products
}
