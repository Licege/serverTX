package messages

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		message := Message{}

		db.QueryRow(`SELECT * FROM messages WHERE id = $1`, id).Scan(
			&message.Id,
			&message.Name,
			&message.Phone,
			&message.Email,
			&message.Content,
			&message.CreateAt)

		c.JSON(http.StatusOK, message)
	}
}
