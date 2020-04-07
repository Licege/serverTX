package files

import "github.com/gin-gonic/gin"

func GetFile(c *gin.Context)  {
	id := c.Param("id")
	dir := DirF + id
	c.File(dir)
}
