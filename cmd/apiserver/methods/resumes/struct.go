package resumes

type Resume = struct {
	Id int `json:"id"`
	CreateAt int64 `json:"create_at"`
	VacancyId int `json:"vacancy_id"`
	Aspirant Aspirant `json:"aspirant"`
	File File `json:"file"`
}

type Aspirant = struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age int `json:"age"`
	Education string `json:"education"`
	Experience int `json:"experience"`
	WorkPlace string `json:"work_place"`
	Phone string `json:"phone"`
	Mail string `json:"mail"`
	Image File `json:"image"`
}

type File = struct {
	Id int `json:"id"`
	Url string `json:"url"`
}