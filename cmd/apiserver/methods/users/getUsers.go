package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryPage := c.DefaultQuery("page", "0")
		perPage := 10
		users := []User{}
		page, _ := strconv.ParseInt(queryPage, 10, 64)
		limit := 10 * page
		var count int

		rows, err := db.Query(`select * from users limit $1 offset $2`, perPage, limit)
		db.QueryRow(`select count(*) from users`).Scan(&count)

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
		result := Result{
			Users: users,
			TotalCount: count,
		}
		c.JSON(http.StatusOK, result)
	}
}
