package documents

type Documents = struct {
	Id int `json:"id"`
	Type int `json:"type"`
	Title string `json:"title"`
	Content string `json:"content"`
	MainTitle bool `json:"main_title"`
}
