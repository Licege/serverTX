package contacts

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContacts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		contacts := Contacts{}

		err := db.QueryRow(`SELECT * FROM contacts`).Scan(
			&contacts.Vk,
			&contacts.Fb,
			&contacts.Tg,
			&contacts.Inst,
			&contacts.Google,
			&contacts.Tw,
			&contacts.Phone,
			&contacts.Schedule,
			&contacts.Address)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, contacts)
	}
}
