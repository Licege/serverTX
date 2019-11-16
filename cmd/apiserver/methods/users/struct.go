package users

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	BonusPoints int `json:"bonus_points"`
	Enabled bool `json:"enabled"`
}

type Users []User

type Result struct {
	Users Users `json:"users"`
	TotalCount int `json:"total_count"`
}

