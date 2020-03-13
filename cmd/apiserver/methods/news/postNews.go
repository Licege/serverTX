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
		err := db.QueryRow(`INSERT INTO news(label, content, create_at, url) VALUES ($1, $2, $3, $4) RETURNING id`,
			requestBody.Label,
			requestBody.Content,
			requestBody.CreateAt,
			requestBody.Url).Scan(&id)

		if err != nil {
			panic(err)
		}

		news := News{
			Id: id,
			Label: requestBody.Label,
			Content: requestBody.Content,
			CreateAt: requestBody.CreateAt,
			Url: requestBody.Url,
		}

		c.JSON(http.StatusOK, news)
	}
}
