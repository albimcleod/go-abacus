package goabacus

//Product defines a product from Abacus POS
type Product struct {
	ProductID   int     `json:"productId"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Taxable     bool    `json:"taxable"`
	Cost        float64 `json:"cost"`
	Category    string  `json:"category"`
	ProductCode string  `json:"productCode"`
}

//Products defines a list of products from Abacus
type Products struct {
	Pagination Pagination `json:"pagination"`
	Products   []Product  `json:"products"`
}
