package news

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test2/cmd/apiserver/methods/files"
)

func GetNews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentPage := c.DefaultQuery("page", "1")
		perPage := 10
		page, _ := strconv.ParseInt(currentPage, 10, 64)
		limit := 10 * (page - 1)

		var newsArr []News
		var totalCount int

		rows, err := db.Query(`SELECT * FROM news LIMIT $1 OFFSET $2`, perPage, limit)
		db.QueryRow(`SELECT count(*) FROM news`).Scan(&totalCount)
		if err != nil {
			panic(err)
		}

		var id, fileId int
		var createAt int64
		var title, description, shortDescription string

		for rows.Next() {
			err := rows.Scan(&id, &title, &description, &createAt, &shortDescription, &fileId)
			if err != nil {
				panic(err)
			}

			newFile := files.File{}
			if fileId != 0 {
				var url string
				err = db.QueryRow(`SELECT url FROM file WHERE id = $1`, fileId).Scan(&url)
				if err != nil {
					c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
					return
				}
				newFile.Id = fileId
				newFile.Url = url
			}

			newNews := News{
				Id:               id,
				Title:            title,
				Description:      description,
				CreateAt:         createAt,
				ShortDescription: shortDescription,
				File:           newFile,
			}
			newsArr = append(newsArr, newNews)
		}
		result := Result{
			NewsList:       newsArr,
			TotalCount: totalCount,
		}
		c.JSON(http.StatusOK, result)
	}
}
