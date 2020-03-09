package orders

type Order = struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Date string `json:"date"`
	Count int `json:"count"`
	Comment string `json:"comment"`
}
