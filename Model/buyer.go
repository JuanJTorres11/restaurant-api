package Model

import (
	"encoding/json"
)

type Buyer struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type BuyerID struct {
	ID string `json:"id,omitempty"`
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
