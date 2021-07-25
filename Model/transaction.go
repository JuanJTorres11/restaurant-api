package Model

type Transaction struct {
	ID       string    `json:"id,omitempty"`
	Buyer    Buyer     `json:"buyer,omitempty"`
	IP       string    `json:"when,omitempty"`
	Device   string    `json:"device,omitempty"`
	Products []Product `json:"products,omitempty"`
}
