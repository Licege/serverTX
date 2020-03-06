package contacts

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutContacts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Contacts{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE contacts SET vk = $1, fb = $2, tg = $3, inst = $4, google = $5, tw = $6, phone = $7, schedule = $8, address = $9`,
			requestBody.Vk,
			requestBody.Fb,
			requestBody.Tg,
			requestBody.Inst,
			requestBody.Google,
			requestBody.Tw,
			requestBody.Phone,
			requestBody.Schedule,
			requestBody.Address)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
