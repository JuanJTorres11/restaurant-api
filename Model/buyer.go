package Model

type Buyer struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
