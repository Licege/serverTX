package staff

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutEmployee (db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := Person{}
		c.BindJSON(&requestBody)

		db.QueryRow(`UPDATE staff SET name = $1, surname = $2, phone = $3, address = $4, profession = $5, file_id = $6 where id = $7`,
			requestBody.Name,
			requestBody.Surname,
			requestBody.Phone,
			requestBody.Address,
			requestBody.Profession,
			requestBody.FileId,
			id)
/*
		if err != nil {
			panic(err)
		}
*/
		c.Status(http.StatusOK)
	}
}
