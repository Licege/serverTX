package reviews

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func GetReviews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var Reviews []Review

		rows, err := db.Query(`SELECT * FROM reviews`)
		if err != nil {
			panic(err.Error())
		}

		var Id, Rating, Status int
		var Forename, Surname, Phone, Description string
		var CreateAt int64
		var RuleAgree bool

		for rows.Next() {
			err := rows.Scan(&Id, &Forename, &Surname, &Phone, &Rating, &Description, &RuleAgree, &CreateAt, &Status)
			if err != nil {
				panic(err.Error())
			}

			rowsFile, err := db.Query(`SELECT id_file FROM reviews_files WHERE id_review = $1`, Id)
			if err != nil {
				panic(err.Error())
			}

			var PhotosId []int
			var PhotoId int
			for rowsFile.Next() {
				err := rowsFile.Scan(&PhotoId)
				if err != nil {
					panic(err.Error())
				}
				PhotosId = append(PhotosId, PhotoId)
			}

			var PhotoArr []File
			var Url string
			if len(PhotosId) > 0 {
				for i := 0; i < len(PhotosId); i++ {
					err := db.QueryRow(`SELECT url FROM file WHERE id = $1`, PhotosId[i]).Scan(&Url)
					if err != nil {
						panic(err.Error())
					}
					Photo := File{
						Id:  PhotosId[i],
						Url: Url,
					}
					PhotoArr = append(PhotoArr, Photo)
				}
			}

			newReview := Review{
				Id:          Id,
				Forename:    Forename,
				Surname:     Surname,
				Phone:       Phone,
				Rating:      Rating,
				Description: Description,
				RuleAgree:   RuleAgree,
				CreateAt:    CreateAt,
				Photo:       PhotoArr,
				Status:      Status,
			}
			Reviews = append(Reviews, newReview)
		}
	}
}
