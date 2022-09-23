package routes

import (
	controller "github.com/itzganesh03/Golang-Practical-/jwtauth/controllers"

	"github.com/itzganesh03/Golang-Practical-/jwtauth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
