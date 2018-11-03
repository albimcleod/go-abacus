package goabacus

//Pagination defines the pagination result for an abacus request
type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
}
