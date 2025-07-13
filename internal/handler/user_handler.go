package handler

import (
	"berbagi/internal/formatter"
	"berbagi/internal/helper"
	"berbagi/internal/models"
	"berbagi/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService services.Service
}

func NewUserHandler(userService services.Service) *userHandler {
	return &userHandler{userService}
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
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	formatter := formatter.FormatUser(newuser, "tokentokentoken")

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
}