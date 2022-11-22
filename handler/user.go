package handler

import (
	"bogistartup/helper"
	"bogistartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}

}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(newUser, "TokenTokenTokenToken")

	response := helper.APIResponse("Account Has Been Registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")
	response := helper.APIResponse("Succesfuly Loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
