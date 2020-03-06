package staff

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostEmployee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Person{}
		c.BindJSON(&requestBody)

		var lastID int
		err := db.QueryRow(`INSERT INTO staff (name, surname, phone, address, id_profession, file_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
			requestBody.Name,
			requestBody.Surname,
			requestBody.Phone,
			requestBody.Address,
			requestBody.Profession,
			requestBody.FileId).Scan(&lastID)

		if err != nil {
			panic(err)
		}

		person := Person{
			Id:         lastID,
			Name:       requestBody.Name,
			Surname:    requestBody.Surname,
			Phone:      requestBody.Phone,
			Address:    requestBody.Address,
			Profession: requestBody.Profession,
			FileId:     requestBody.FileId,
		}


		c.JSON(http.StatusOK, person)
	}
}
