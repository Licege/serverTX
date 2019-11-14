package vacancies

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteVacancy(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec(`DELETE FROM vacancy WHERE id = $1`, id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
