package staff

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEmployeeById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		person := Person{}

		db.QueryRow(`select * from staff where id = $1`, id).Scan(
			&person.Id,
			&person.Name,
			&person.Surname,
			&person.Phone,
			&person.Address,
			&person.Profession,
			&person.FileId)

		c.JSON(http.StatusOK, person)
	}
}
