package news

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test2/cmd/apiserver/methods/files"
)

func GetNewsById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var newsId, fileId int
		var createAt int64
		var title, description, shortDesctiption string
		db.QueryRow(`SELECT * FROM news WHERE id = $1`, id).Scan(
			&newsId,
			&title,
			&description,
			&createAt,
			&shortDesctiption,
			&fileId)

		newFile := files.File{}
		if fileId != 0 {
			var url string
			err := db.QueryRow(`SELECT url FROM file WHERE id = $1`, fileId).Scan(&url)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
				return
			}
			newFile.Id = fileId
			newFile.Url = url
		}

		news := News{
			Id:               newsId,
			Title:            title,
			Description:      description,
			CreateAt:         createAt,
			ShortDescription: shortDesctiption,
			File:             newFile,
		}

		c.JSON(http.StatusOK, news)
	}
}
