package categories

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := Category{}
		c.BindJSON(&requestBody)

		err := db.QueryRow(`UPDATE categories SET title = $1, title_en = $2 WHERE id = $3`,
			requestBody.Title,
			requestBody.TitleEn,
			id)
		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
