package staff

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteEmployee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM staff WHERE id = $1", id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
