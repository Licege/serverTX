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
		var id, weight, price, category int
		var title, description, url string

		for rows.Next(){
			err := rows.Scan(&id, &title, &description, &category, &price, &weight, &url)
			if err != nil {
				panic(err)
			}
			newDish := Dish{
				Id:          id,
				Title:       title,
				Description: description,
				CategoryId:    category,
				Price:       price,
				Weight:      weight,
				Url:      url,
			}
			dishes = append(dishes, newDish)
		}
		c.JSON(http.StatusOK, dishes)
	}
}
