package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}

}

func (h *userHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "user_index.html", nil)

}
