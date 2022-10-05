package handler

import (
	"learnGo/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	//Tangkap input dari user
	// map input dari user ke struck RegisterUserInput
	// struck di atas dipassing sebagai parameter service

	var input user.RegisterUserInput

	err != c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)

	}

}
