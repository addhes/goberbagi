package handler

import (
	"berbagi/internal/auth"
	"berbagi/internal/formatter"
	"berbagi/internal/helper"
	"berbagi/internal/models"
	"berbagi/internal/services"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService services.Service
	authService auth.Service
}

func NewUserHandler(userService services.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input models.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, nil)
	// }

	//     err := c.ShouldBindJSON(&input)
	if err != nil {

		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}

		// response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		// c.JSON(http.StatusUnprocessableEntity, response)
		// return

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}

			response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		// Jika error bukan validasi, misalnya JSON tidak bisa di-parse
		response := helper.APIResponse("Invalid input format", http.StatusBadRequest, "error", gin.H{"errors": err.Error()})
		c.JSON(http.StatusBadRequest, response)
		return

	}

	newuser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authService.GenerateToken(newuser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	formatter := formatter.FormatUser(newuser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "sucess", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input email and password
	// input ditangkap handler
	// maping dari input user ke input struct
	// input struct passing service
	// di service mencari dg bantuan repository user dengan email x
	// mencocoan password

	var input models.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	formatter := formatter.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input dari user
	// input email dimapping ke struct input
	// struct input dipassing ke service
	// service akan memanggil repository - email sudah ada atau belum
	// repository - db

	var input models.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"Errors": errors}

		response := helper.APIResponse("failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email cheching failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// simpan gambarnya difolder "images/"
	// diservice kita panggil repo
	// JWT (sementara hardcode , seakan2x user yang login ID = 1)
	// repo ambil data user yang ID = 1
	// repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		// Gagal ambil file dari form
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Harusnya dapat jwt,
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		// Gagal simpan file
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		// Gagal simpan file
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
