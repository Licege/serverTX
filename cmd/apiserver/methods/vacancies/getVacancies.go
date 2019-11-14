package vacancies

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVacancies(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		vacancies := []Vacancy{}

		rows, err := db.Query(`SELECT * FROM vacancy`)

		if err != nil {
			panic(err)
		}

		var id, fileId, salaryFrom, salaryTo int
		var title, requirements, description string
		for rows.Next() {
			err := rows.Scan(&id, &title, &requirements, &description, &salaryFrom, &salaryTo, &fileId)

			if err != nil {
				panic(err)
			}

			newVacancy := Vacancy{
				Id:           id,
				Title:        title,
				Requirements: requirements,
				Description:  description,
				SalaryFrom:   salaryFrom,
				SalaryTo:     salaryTo,
				FileId:       fileId,
			}
			vacancies = append(vacancies, newVacancy)
		}
		c.JSON(http.StatusOK, vacancies)
	}
}
