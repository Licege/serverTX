package cities

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutCities(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := City{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE cities SET title = $1 WHERE id = $2`,
			&requestBody.Title,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
