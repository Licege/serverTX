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

		db.QueryRow(`UPDATE professions SET profession = $1 WHERE id = $2`,
			requestBody.Profession,
			id)

		c.Status(http.StatusOK)
	}
}
