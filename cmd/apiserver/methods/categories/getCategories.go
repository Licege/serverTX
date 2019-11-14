package categories

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := []Category{}

		rows, err := db.Query(`SELECT * FROM categories`)

		if err != nil {
			panic(err)
		}

		var id int
		var title string
		for rows.Next() {
			err := rows.Scan(&id, &title)
			if err != nil {
				panic(err)
			}

			newCategory := Category{
				Id:    id,
				Title: title,
			}
			categories = append(categories, newCategory)
		}
		c.JSON(http.StatusOK, categories)
	}
}
