package staff

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEmployees(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		persons := []Person{}

		rows, err := db.Query(`select * from staff`)

		if err != nil {
			panic(err)
		}

		var id, profession, fileId int
		var name, surname, phone, address string
		for rows.Next() {
			err := rows.Scan(&id, &name, &surname, &phone, &address, &profession, &fileId)

			if err != nil {
				panic(err)
			}

			newPerson := Person{
				Id:         id,
				Name:       name,
				Surname:    surname,
				Phone:      phone,
				Address:    address,
				Profession: profession,
				FileId:     fileId,
			}
			persons = append(persons, newPerson)
		}
		c.JSON(http.StatusOK, persons)
	}
}
