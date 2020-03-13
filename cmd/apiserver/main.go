package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"test2/cmd/apiserver/methods/categories"
	"test2/cmd/apiserver/methods/contacts"
	"test2/cmd/apiserver/methods/menu"
	"test2/cmd/apiserver/methods/messages"
	"test2/cmd/apiserver/methods/news"
	"test2/cmd/apiserver/methods/orders"
	"test2/cmd/apiserver/methods/profession"
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

	// вырубаем CORS
	//r.Use(LiberalCORS)
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	//r.Static("/stat-img", "./image")

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
	professionR := r.Group("/api/professions")
	{
		professionR.GET("/", profession.GetProfessions(db))
		professionR.POST("/", profession.PostProfession(db))
		professionR.PUT("/:id", profession.PutProfession(db))
	}
	contactsR := r.Group("/api/contacts")
	{
		contactsR.GET("/", contacts.GetContacts(db))
		contactsR.PUT("/", contacts.PutContacts(db))
	}
	newsR := r.Group("/api/news")
	{
		newsR.GET("/", news.GetNews(db))
		newsR.GET("/:id", news.GetNewsById(db))
		newsR.POST("/", news.PostNews(db))
		newsR.PUT("/:id", news.UpdateNews(db))
		newsR.DELETE("/:id", news.DeleteNews(db))
	}
	ordersR := r.Group("/api/orders")
	{
		ordersR.POST("/", orders.PostOrder(db))
		ordersR.GET("/", orders.GetOrders(db))
	}
	messagesR := r.Group("/api/messages")
	{
		messagesR.GET("/", messages.GetMessages(db))
		messagesR.GET("/:id", messages.GetMessage(db))
		messagesR.DELETE("/:id", messages.DeleteMessage(db))
	}


	r.Run(":9090")

	/*

	func LiberalCORS(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
				c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
			}
			c.AbortWithStatus(http.StatusOK)
		}
	}
	 */
}