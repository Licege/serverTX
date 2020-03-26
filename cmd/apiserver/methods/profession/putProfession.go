package profession

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutProfession (db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := Profession{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE professions SET profession = $1 WHERE id = $2`,
			requestBody.Profession,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusOK)
	}
}
