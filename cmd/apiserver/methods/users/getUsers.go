package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users := []User{}

		rows, err := db.Query(`select * from users`)

		if err != nil {
			panic(err)
		}

		var id, bonusPoints int
		var email, phone, name, surname string
		var enabled bool
		for rows.Next() {
			err := rows.Scan(&id, &email, &phone, &name, &surname, &bonusPoints, &enabled)

			if err != nil {
				panic(err)
			}

			newUser := User{
				Id: id,
				Name: name,
				Surname: surname,
				Email: email,
				Phone: phone,
				BonusPoints: bonusPoints,
				Enabled: enabled,
			}
			users = append(users, newUser)
		}
		c.JSON(http.StatusOK, users)
	}
}
