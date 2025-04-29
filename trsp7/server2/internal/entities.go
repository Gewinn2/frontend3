package internal

type Product struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category []string `json:"category"`
}

type ProductWithoutId struct {
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category []string `json:"category"`
}

type CreateProductRequest struct {
	Products []ProductWithoutId `json:"products"`
}

type Catalog struct {
	Products []Product `json:"products"`
}

type Error struct {
	Error string `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}

type Ids struct {
	Ids []int `json:"ids"`
}
