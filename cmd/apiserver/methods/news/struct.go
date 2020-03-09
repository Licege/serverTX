package news

type News = struct {
	Id int `json:"id"`
	Label string `json:"label"`
	Content string `json:"content"`
	CreateAt int `json:"create_at"`
	Url string `json:"url"`
}
