package file

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func MyUploadImage(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")

		filename := filepath.Base(file.Filename)
		err := c.SaveUploadedFile(file, filename)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка: %s", err.Error()))
			return
		}

		c.Status(http.StatusOK)
	}
}
