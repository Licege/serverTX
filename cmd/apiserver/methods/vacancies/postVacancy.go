package vacancies

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostVacancy(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Vacancy{}
		c.BindJSON(&requestBody)

		var lastID int
		err := db.QueryRow(`INSERT INTO vacancy(title, requirements, description, salary_from, salary_to, file_id)`,
			requestBody.Title,
			requestBody.Requirements,
			requestBody.Description,
			requestBody.SalaryFrom,
			requestBody.SalaryTo,
			requestBody.FileId).Scan(&lastID)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, lastID)
	}
}
