package cities

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCities(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cities := []City{}

		rows, err := db.Query(`SELECT * FROM cities`)

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

			city := City{
				Id:    id,
				Title: title,
			}

			cities = append(cities, city)
		}
		c.JSON(http.StatusOK, cities)
	}
}
