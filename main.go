package main

import (
	"berbagi/internal/handler"
	"berbagi/internal/repository"
	"berbagi/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// fmt.Println("Hello, World")

	  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
	
	dsn := "root:@tcp(127.0.0.1:3306)/db_go_berbagi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewRepository(db)
	userService := services.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run();

	// userInput := models.RegisterUserInput{}
	// userInput.Name = "Awan"
	// userInput.Email = "awan@g.com"
	// userInput.Occupation = "Waiters"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)
}

// func handler(c *gin.Context) {

// 	var users []models.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }