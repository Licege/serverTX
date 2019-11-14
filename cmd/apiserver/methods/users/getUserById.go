package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsersByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user := User{}

		err := db.QueryRow(`select * from users where id = $1`, id).Scan(
			&user.Id,
			&user.Email,
			&user.Phone,
			&user.Name,
			&user.Surname,
			&user.BonusPoints,
			&user.Enabled)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, user)
	}
}
