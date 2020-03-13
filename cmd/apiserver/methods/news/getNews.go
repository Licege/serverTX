package news

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsArr := []News{}

		rows, err := db.Query(`SELECT * FROM news`)
		if err != nil {
			panic(err)
		}

		var id int
		var create_at int64
		var label, content, url string

		for rows.Next() {
			err := rows.Scan(&id, &label, &content, &create_at, &url)
			if err != nil {
				panic(err)
			}
			newNews := News {
				Id: id,
				Label: label,
				Content: content,
				CreateAt: create_at,
				Url: url,
			}
			newsArr = append(newsArr, newNews)
		}
		c.JSON(http.StatusOK, newsArr)
	}
}
