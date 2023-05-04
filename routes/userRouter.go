package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shittu24/jwt/controllers"
	"github.com/shittu24/jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
