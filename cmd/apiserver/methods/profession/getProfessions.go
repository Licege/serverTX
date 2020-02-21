package profession

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfessions (db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		professions := []Profession{}

		rows, err := db.Query(`select * from profession`)

		if err != nil {
			panic(err)
		}

		var id int
		var profession string
		for rows.Next() {
			err := rows.Scan(&id, &profession)

			if err != nil {
				panic(err)
			}

			newProfession := Profession{
				Id: id,
				Profession: profession,
			}
			professions = append(professions, newProfession)
		}
		c.JSON(http.StatusOK, professions)
	}
}
