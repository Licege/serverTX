package news

import "test2/cmd/apiserver/methods/files"

type News = struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreateAt int64 `json:"create_at"`
	ShortDescription string `json:"short_description"`
	File files.File `json:"file"`
}

type Result = struct {
	NewsList []News `json:"news_list"`
	TotalCount int `json:"total_count"`
}
