package menu

type Dish struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Weight int `json:"weight"`
	Price int `json:"price"`
	Category int `json:"category"`
	FileId int `json:"file_id"`
}
