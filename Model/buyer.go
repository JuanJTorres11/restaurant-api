package Model

import (
	"encoding/json"
)

type Buyer struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type SimpleBuyer struct {
	Name         string              `json:"name,omitempty"`
	Age          int                 `json:"age,omitempty"`
	Transactions []SimpleTransaction `json:"transactions,omitempty"`
}

type BuyerID struct {
	ID string `json:"id,omitempty"`
}

type BuyerName struct {
	Name string `json:"name,omitempty"`
}

func formatBuyers(r []byte) []Buyer {
	var buyers []Buyer
	json.Unmarshal(r, &buyers)
	allKeys := make(map[string]bool)
	list := []Buyer{}
	for _, item := range buyers {
		if _, value := allKeys[item.ID]; !value {
			allKeys[item.ID] = true
			list = append(list, item)
		}
	}

	return list
}
