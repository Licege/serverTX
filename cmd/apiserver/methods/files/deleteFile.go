package files

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func DeleteFile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var url string
		err := db.QueryRow(`DELETE FROM file WHERE id = $1 RETURNING url`, id).Scan(&url)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
			return
		}

		err = os.Remove(url)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
			return
		}

		c.Status(http.StatusNoContent)
	}
}
