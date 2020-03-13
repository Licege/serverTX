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

		_, err := db.Exec(`UPDATE news SET label = $1, content = $2, url = $3 WHERE id = $4`,
			requestBody.Label,
			requestBody.Content,
			requestBody.Url,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
