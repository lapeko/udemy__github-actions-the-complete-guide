package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lapeko/udemy__github-actions-the-complete-guide/section3-go-api/internal/api/controller"
)

func SetupUsers(r *gin.RouterGroup) {
	r.POST("/", controller.PostUserController)
	r.GET("/", controller.GetUsersController)
	r.GET("/:id", controller.GetUserByIdController)
	r.PUT("/:id", controller.PutUserController)
	r.DELETE("/:id", controller.DeleteUserController)
}
