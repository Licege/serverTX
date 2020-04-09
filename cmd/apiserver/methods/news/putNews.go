package news

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateNews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := News{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE news SET title = $1, description = $2, short_description = $3, file_id = $4 WHERE id = $5`,
			requestBody.Title,
			requestBody.Description,
			requestBody.ShortDescription,
			requestBody.File.Id,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
