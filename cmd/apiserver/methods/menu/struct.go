package menu

type Dish struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Weight int `json:"weight"`
	Price int `json:"price"`
	CategoryId int `json:"category_id"`
	Url string `json:"url"`
}
