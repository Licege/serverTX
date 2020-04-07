package reviews

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostReview(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Review{}
		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err.Error())
		}

		var Id int
		err = db.QueryRow(`INSERT INTO reviews (forename, surname, phone, rating, description, rule_agree, create_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
			requestBody.Forename,
			requestBody.Surname,
			requestBody.Phone,
			requestBody.Rating,
			requestBody.Description,
			requestBody.RuleAgree,
			requestBody.CreateAt).Scan(&Id)

		if err !=nil {
			panic(err.Error())
		}

		for i := 0; i < len(requestBody.Photo); i++ {
			_, err = db.Exec(`INSERT INTO reviews_files(id_review, id_file) VALUES ($1, $2)`,
				Id,
				requestBody.Photo[i].Id)

			if err != nil {
				panic(err.Error())
			}
		}

		c.Status(http.StatusOK)
	}
}
