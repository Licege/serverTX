package news

type News = struct {
	Id int `json:"id"`
	Label string `json:"label"`
	Content string `json:"content"`
	CreateAt int64 `json:"create_at"`
	Url string `json:"url"`
}
