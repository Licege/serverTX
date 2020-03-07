package categories

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Category{}
		c.BindJSON(&requestBody)

		var lastID int
		err := db.QueryRow(`INSERT INTO categories(title, title_en) VALUES ($1, $2) RETURNING id`,
			requestBody.Title,
			requestBody.TitleEn).Scan(&lastID)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, lastID)
	}
}
