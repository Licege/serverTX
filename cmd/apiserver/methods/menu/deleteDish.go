package menu

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteDish(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM dishes WHERE id = $1", id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
