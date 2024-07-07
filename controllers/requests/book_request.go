package requests

type BookRequest struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
