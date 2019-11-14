package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"test2/cmd/apiserver/methods/categories"
	"test2/cmd/apiserver/methods/menu"
	"test2/cmd/apiserver/methods/staff"
	"test2/cmd/apiserver/methods/users"
	"test2/cmd/apiserver/methods/vacancies"
)

var (
	db *sql.DB
)

func main()  {
	var err error

	connStr := "user=postgres password=5315 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	usersR := r.Group("/api/users")
	{
		usersR.GET("/", users.GetUsers(db))
		usersR.POST("/", users.CreateUser(db))
		usersR.GET("/:id", users.GetUsersByID(db))
		usersR.PUT("/:id", users.PutUser(db))
	}
	employeeR := r.Group("/api/employees")
	{
		employeeR.GET("/", staff.GetEmployees(db))
		employeeR.GET("/:id", staff.GetEmployeeById(db))
		employeeR.POST("/", staff.PostEmployee(db))  // Инты отправлять не в ""
		employeeR.PUT("/:id", staff.PutEmployee(db))
		employeeR.DELETE("/:id", staff.DeleteEmployee(db))
	}
	menuR := r.Group("/api/menu")
	{
		menuR.GET("/", menu.GetDishes(db))
		menuR.GET("/:id", menu.GetDishById(db))
		menuR.POST("/", menu.PostDish(db))
		menuR.PUT("/:id", menu.PutDish(db))
		menuR.DELETE("/:id", menu.DeleteDish(db))
	}
	categoryR := r.Group("/api/categories")
	{
		categoryR.GET("/", categories.GetCategories(db))
		categoryR.POST("/", categories.PostCategory(db))
		categoryR.PUT("/:id", categories.PutCategory(db))
		categoryR.DELETE("/:id", categories.DeleteCategory(db))
	}
	vacancyR := r.Group("/api/vacancy")
	{
		vacancyR.GET("/", vacancies.GetVacancies(db))
		vacancyR.GET("/:id", vacancies.GetVacancyById(db))
		vacancyR.POST("/", vacancies.PostVacancy(db))
		vacancyR.PUT("/:id", vacancies.PutVacancy(db))
		vacancyR.DELETE("/:id", vacancies.DeleteVacancy(db))
	}

	r.Run(":9090")
}