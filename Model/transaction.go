package Model

import (
	"strings"
)

type Transaction struct {
	ID       string      `json:"id,omitempty"`
	Buyer    BuyerID     `json:"buyer,omitempty"`
	IP       string      `json:"ip,omitempty"`
	Device   string      `json:"device,omitempty"`
	Products []ProductID `json:"products,omitempty"`
}

func formatTransactions(r []byte) []Transaction {
	var transactions []Transaction
	p := string(r)
	replacer := strings.NewReplacer("#", "", "(", "", ")", "")
	p = replacer.Replace(p)

	result := strings.Split(p, "\x00\x00")

	allKeys := make(map[string]bool)
	for _, item := range result {
		t := strings.Split(item, "\x00")
		if _, value := allKeys[t[0]]; !value && len(t) > 1 {
			allKeys[t[0]] = true
			products := strings.Split(t[4], ",")
			var list []ProductID
			for _, v := range products {
				list = append(list, ProductID{v})
			}
			transactions = append(transactions, Transaction{t[0], BuyerID{t[1]}, t[2], t[3], list})
		}
	}

	return transactions
}
