package menu

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDishes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dishes := []Dish{}

		rows, _ := db.Query(`SELECT * FROM dishes`)
		var id, weight, price, category, fileId int
		var title, description string

		for rows.Next(){
			err := rows.Scan(&id, &title, &description, &weight, &price, &category, &fileId)
			if err != nil {
				panic(err)
			}
			newDish := Dish{
				Id:          id,
				Title:       title,
				Description: description,
				Weight:      weight,
				Price:       price,
				Category:    category,
				FileId:      fileId,
			}
			dishes = append(dishes, newDish)
		}
		c.JSON(http.StatusOK, dishes)
	}
}
