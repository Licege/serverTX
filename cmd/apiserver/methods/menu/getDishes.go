package menu

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test2/cmd/apiserver/methods/files"
)

func GetDishes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dishes []Dish

		rows, _ := db.Query(`SELECT * FROM dishes`)
		var id, weight, price, categoryId, fileId int
		var title, description string

		for rows.Next(){
			err := rows.Scan(&id, &title, &description, &categoryId, &price, &weight, &fileId)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
				return
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


			newDish := Dish{
				Id:          id,
				Title:       title,
				Description: description,
				CategoryId:    categoryId,
				Price:       price,
				Weight:      weight,
				File: newFile,
			}
			dishes = append(dishes, newDish)
		}
		c.JSON(http.StatusOK, dishes)
	}
}
