package reviews

type Review = struct {
	Id int `json:"id"`
	Forename string `json:"forename"`
	Surname string `json:"surname"`
	Phone string `json:"phone"`
	Rating int `json:"rating"`
	Description string `json:"description"`
	RuleAgree bool `json:"rule_agree"`
	CreateAt int64 `json:"create_at"`
	Photo []File `json:"photo"`
	Status int `json:"status"`
}

type File = struct {
	Id int `json:"id"`
	Url string `json:"url"`
}
