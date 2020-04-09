package news

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostNews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := News{}
		c.BindJSON(&requestBody)

		var id int
		err := db.QueryRow(`INSERT INTO news(title, description, create_at, short_description, file_id) VALUES ($1, $2, $3, $4) RETURNING id`,
			requestBody.Title,
			requestBody.Description,
			requestBody.CreateAt,
			requestBody.ShortDescription,
			requestBody.File.Id).Scan(&id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, id)
	}
}
