package staff

type Person struct {
	Id         int `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Phone string `json:"phone"`
	Address    string `json:"address"`
	Profession int `json:"profession"`
	FileId     int `json:"file_id"`
}