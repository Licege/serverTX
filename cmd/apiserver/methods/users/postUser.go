package users

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := User{}
		c.BindJSON(&requestBody)
		fmt.Println(requestBody)

		var lastID int
		db.QueryRow("INSERT INTO users(name, surname, email, phone, bonus_points, enabled) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			requestBody.Name, requestBody.Surname, requestBody.Email, requestBody.Phone, 0, true).Scan(&lastID)

		c.JSON(http.StatusOK, lastID)
	}
}
