package news

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNewsById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		news := News{}

		db.QueryRow(`SELECT * FROM news WHERE id = $1`, id).Scan(
			&news.Id,
			&news.Label,
			&news.Content,
			&news.CreateAt,
			&news.Url)

		c.JSON(http.StatusOK, news)
	}
}
