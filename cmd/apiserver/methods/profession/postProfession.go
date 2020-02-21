package profession

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostProfession(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Profession{}
		c.BindJSON(&requestBody)

		var Id int
		err := db.QueryRow(`INSERT INTO professions(profession) VALUES ($1) RETURNING id`,
			requestBody.Profession).Scan(&Id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, Id)
	}
}
