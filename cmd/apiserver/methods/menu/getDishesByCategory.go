package menu

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test2/cmd/apiserver/methods/files"
)

func GetDishesByCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Param("category")

		var categoryId int
		err := db.QueryRow(`SELECT id FROM categories WHERE title_en = $1`, category).Scan(&categoryId)
		if err != nil {
			c.String(http.StatusBadRequest, "Ошибка: %s", err.Error())
		}

		var dishes []Dish
		var id, weight, price, fileId int
		var title, description string
		rows, err := db.Query(`SELECT * FROM dishes WHERE category_id = $1`, categoryId)
		if err != nil {
			c.String(http.StatusNoContent, "По запросу ничего не найдено")
			return
		}
		for rows.Next() {
			err = rows.Scan(&id, &title, &description, &categoryId, &price, &weight, &fileId)
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
		if dishes == nil {
			c.Status(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, dishes)
	}
}
