package entities

type Product struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category []string `json:"category"`
}
