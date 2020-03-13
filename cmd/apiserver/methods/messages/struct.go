package messages

type Message = struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Content string `json:"content"`
	CreateAt int64 `json:"create_at"`
}
