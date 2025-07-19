package main

import (
	"berbagi/internal/auth"
	"berbagi/internal/handler"
	"berbagi/internal/helper"
	"berbagi/internal/repository"
	"berbagi/internal/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

	// userInput := models.RegisterUserInput{}
	// userInput.Name = "Awan"
	// userInput.Email = "awan@g.com"
	// userInput.Occupation = "Waiters"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)
}

func authMiddleware(authService auth.Service, userService services.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer tokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token , err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

// func handler(c *gin.Context) {

// 	var users []models.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
