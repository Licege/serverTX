package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := User{}
		errBind := c.BindJSON(&requestBody)

		if errBind != nil {
			panic(errBind)
		}

		_, err := db.Exec(`update users set email = $1, phone = $2, name = $3, surname = $4, bonus_points = $5, enabled = $6 where id = $7`,
			requestBody.Email, requestBody.Phone, requestBody.Name, requestBody.Surname, requestBody.BonusPoints, requestBody.Enabled, id,
		)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, requestBody)
		//c.Status(http.StatusNoContent)
	}
}
