package orders

type Order = struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Date int64 `json:"date"`
	Count int `json:"count"`
	Comment string `json:"comment"`
}
