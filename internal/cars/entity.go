package cars

type Car struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
