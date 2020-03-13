package messages

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessages(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		messages := []Message{}

		rows, _ := db.Query(`SELECT * FROM messages`)

		var id int
		var createAt int64
		var name, phone, email, content string

		for rows.Next() {
			err := rows.Scan(&id, &name, &phone, &content, &createAt, &email)
			if err != nil {
				panic(err)
			}
			message := Message{
				Id:       id,
				Name:     name,
				Phone:    phone,
				Content:  content,
				CreateAt: createAt,
				Email: 	email,
			}
			messages = append(messages, message)
		}

		c.JSON(http.StatusOK, messages)
	}
}