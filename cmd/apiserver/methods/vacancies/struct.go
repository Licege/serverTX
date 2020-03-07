package vacancies

type Vacancy struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Requirements string `json:"requirements"`
	Description string `json:"description"`
	SalaryFrom int `json:"salary_from"`
	SalaryTo int `json:"salary_to"`
	Url string `json:"url"`
}
