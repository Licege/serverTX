package vacancies

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutVacancy(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := Vacancy{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE vacancy SET title = $1, requirements = $2, description = $3, salary_from = $4, salary_to = $5, url = $6 WHERE id = $7`,
			requestBody.Title,
			requestBody.Requirements,
			requestBody.Description,
			requestBody.SalaryFrom,
			requestBody.SalaryTo,
			requestBody.Url,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
