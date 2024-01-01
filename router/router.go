package router

import (
	"/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.LoginUser)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.DeleteUser)

	// Photo routes
	router.POST("/photos", controllers.AddPhoto)
	router.GET("/photos", controllers.GetAllPhotos)
	router.PUT("/photos/:photoId", controllers.UpdatePhoto)
	router.DELETE("/photos/:photoId", controllers.DeletePhoto)

	return router
}
