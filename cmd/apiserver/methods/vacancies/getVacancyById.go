package vacancies

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVacancyById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		vacancy := Vacancy{}

		err := db.QueryRow(`SELECT * FROM vacancy WHERE id = $1`, id).Scan(
			&vacancy.Id,
			&vacancy.Title,
			&vacancy.Requirements,
			&vacancy.Description,
			&vacancy.SalaryFrom,
			&vacancy.SalaryTo,
			&vacancy.FileId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, vacancy)
	}
}
