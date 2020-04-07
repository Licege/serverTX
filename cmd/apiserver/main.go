package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"test2/cmd/apiserver/methods/categories"
	"test2/cmd/apiserver/methods/cities"
	"test2/cmd/apiserver/methods/contacts"
	"test2/cmd/apiserver/methods/delivery"
	"test2/cmd/apiserver/methods/delivery/delivery_global_settings"
	"test2/cmd/apiserver/methods/delivery/delivery_settings"
	"test2/cmd/apiserver/methods/files"
	"test2/cmd/apiserver/methods/menu"
	"test2/cmd/apiserver/methods/messages"
	"test2/cmd/apiserver/methods/news"
	"test2/cmd/apiserver/methods/orders"
	"test2/cmd/apiserver/methods/profession"
	"test2/cmd/apiserver/methods/reviews"
	"test2/cmd/apiserver/methods/staff"
	"test2/cmd/apiserver/methods/users"
	"test2/cmd/apiserver/methods/vacancies"
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
	r.MaxMultipartMemory = 10 << 20 //10 MiB

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
		employeeR.POST("/", staff.PostEmployee(db))
		employeeR.PUT("/:id", staff.PutEmployee(db))
		employeeR.DELETE("/:id", staff.DeleteEmployee(db))
	}
	menuR := r.Group("/api/menu")
	{
		menuR.GET("/", menu.GetDishes(db))
		menuR.GET("/dish/:id", menu.GetDishById(db))
		menuR.GET("/category/:category", menu.GetDishesByCategory(db))
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
	deliveryR := r.Group("/api/delivery")
	{
		deliveryR.GET("/", delivery.GetDeliveryOrders(db))
		deliveryR.POST("/", delivery.PostDelivery(db))
	}
	deliverySettingsR := r.Group("/api/delivery/settings")
	{
		deliverySettingsR.GET("/", delivery_settings.GetDeliverySettings(db))
	}
	deliveryGlobalSettingsR := r.Group("/api/delivery/global-settings")
	{
		deliveryGlobalSettingsR.GET("/", delivery_global_settings.GetDeliveryGlobalSettings(db))
		deliveryGlobalSettingsR.PUT("/", delivery_global_settings.PutDeliveryGlobalSettings(db))
	}
	cityR := r.Group("/api/cities")
	{
		cityR.GET("/", cities.GetCities(db))
		cityR.PUT("/:id", cities.PutCities(db))
	}
	reviewsR := r.Group("/api/reviews")
	{
		reviewsR.GET("/", reviews.GetReviews(db))
		reviewsR.POST("/", reviews.PostReview(db))
	}

	imageR := r.Group("/api/file")
	{
		imageR.POST("/", files.FileUploadFunc(db))
		imageR.DELETE("/:id", files.DeleteFile(db))
	}

	r.GET("/api/upload/:id", files.GetFile) // получаем файл по ссылке


	r.Run(":9090")
}