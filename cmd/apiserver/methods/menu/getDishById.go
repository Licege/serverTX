package menu

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test2/cmd/apiserver/methods/files"
)

func GetDishById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var dishId, weight, price, categoryId, fileId int
		var title, description, url string
		db.QueryRow(`SELECT * FROM dishes WHERE id = $1`, id).Scan(
			&dishId,
			&title,
			&description,
			&categoryId,
			&price,
			&weight,
			&fileId)

		newFile := files.File{}
		if fileId != 0 {
			err := db.QueryRow(`SELECT url FROM file WHERE id = $1`, fileId).Scan(&url)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
				return
			}
			newFile.Id = fileId
			newFile.Url = url
		}

		dish := Dish{
			Id:          dishId,
			Title:       title,
			Description: description,
			Weight:      weight,
			Price:       price,
			CategoryId:  categoryId,
			File:        newFile,
		}

		c.JSON(http.StatusOK, dish)
	}
}
