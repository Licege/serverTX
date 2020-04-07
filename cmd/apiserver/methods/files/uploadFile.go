package files

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

func FileUploadFunc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка1: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		arr := strings.Split(filename, ".")
		var fileExt string
		if len(arr) > 1 {
			fileExt = arr[len(arr) - 1]
		}
		newFilename := fmt.Sprintf("%s.%s", randomFileName(), fileExt)

		dir := filepath.Join(DirF, newFilename)

		err = c.SaveUploadedFile(file, dir)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка2: %s", err.Error()))
			return
		}

		url := APIUrl + newFilename
		var fileId int
		err = db.QueryRow(`INSERT INTO file(url) VALUES ($1) RETURNING id`, url).Scan(&fileId)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка3: %s", err.Error()))
			return
		}

		data := File{
			Id:  fileId,
			Url: url,
		}

		c.JSON(http.StatusOK, data)
	}
}
