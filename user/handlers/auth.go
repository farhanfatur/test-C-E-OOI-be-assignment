package handlers

import (
	"net/http"

	"github.com/farhanfatur/assignment-be/user/domains"
	"github.com/farhanfatur/assignment-be/user/usecases"
	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	gin     *gin.Engine
	usecase *usecases.AuthUsecase
}

func NewAuthHandlers(gin *gin.Engine, usecase *usecases.AuthUsecase) *AuthHandlers {
	handler := &AuthHandlers{
		gin:     gin,
		usecase: usecase,
	}

	gin.POST("/auth/login", handler.Login)
	gin.POST("/auth/register", handler.Register)
	gin.POST("/auth/logout", handler.Logout)

	return handler
}

func (a *AuthHandlers) Login(c *gin.Context) {
	body := domains.LoginRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tokenString, err := a.usecase.Attempt(c, body.Username, body.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(400, domains.ResponseJSON{
			Data: "Invalid Error",
		})
		return
	}

	c.JSON(200, domains.ResponseJSON{
		Data:    tokenString,
		Message: "Login is success",
	})
}

func (a *AuthHandlers) Register(c *gin.Context) {
	body := domains.RegisterRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := a.usecase.Register(body.Username, body.Password, body.AccountType)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(400, domains.ResponseJSON{
			Data: "Invalid Error",
		})
		return
	}

	c.JSON(200, domains.ResponseJSON{
		Data:    "",
		Message: "Register is success",
	})
}

func (a *AuthHandlers) Logout(c *gin.Context) {
	token, _ := c.Cookie("token_set")
	if err := a.usecase.Logout(token); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(400, domains.ResponseJSON{
			Data: err,
		})
		return
	}
	c.JSON(200, domains.ResponseJSON{
		Data:    "",
		Message: "Logout is success",
	})
}
