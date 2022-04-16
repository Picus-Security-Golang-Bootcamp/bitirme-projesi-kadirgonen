package handler

import (
	"net/http"

	model "HW/app/models"
	"HW/app/service"
	"HW/pkg/logger"

	"github.com/gin-gonic/gin"
)

type (
	User struct {
		userService service.UserService
		authService service.JWTAuthService
		logger      logger.Logger
	}

	signupRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	authResponse struct {
		Token string `json:"token"`
	}
	response struct {
		Message string `json:"message" example:"message"`
	}
)

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

func successResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, response{msg})
}
func NewUser(us service.UserService, as service.JWTAuthService, l logger.Logger) *User {
	return &User{us, as, l}
}
func (c *Auth) SignUp(g *gin.Context) {
	var req signupRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "signup")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}
	user := model.NewUser(req.Email, req.Password, []*model.Role{{Name: "customer"}})
	err := c.userService.CreateUser(user)
	if err != nil {
		c.logger.Error(err, "signup")
		errorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := c.authService.CreateToken(*user)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	g.JSON(http.StatusOK, authResponse{Token: *token})
}

func (c *Auth) Login(g *gin.Context) {
	var req loginRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "login")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	user := c.userService.GetUser(req.Email, req.Password)
	if user == nil {
		errorResponse(g, http.StatusNotFound, "invalid email or password")
		return
	}

	token, err := c.authService.CreateToken(*user)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	g.JSON(http.StatusOK, authResponse{Token: *token})
}
