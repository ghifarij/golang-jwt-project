package routes

import (
	controller "github.com/ghifarij/golang-jwt-project/controllers"
	"github.com/ghifarij/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRequest *gin.Engine) {
	incomingRequest.Use(middleware.Authenticate())
	incomingRequest.GET("/users", controller.GetUsers())
	incomingRequest.GET("/users/:user_id", controller.GetUserByID())
}
