package routes

import (
	v1 "University-Information-Website/api/v1"
	"University-Information-Website/middleware"
	"University-Information-Website/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("user/:id", v1.UpdateUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		auth.POST("user/add", v1.AddUser)
		auth.GET("users", v1.GetUsers)
	}
	//router := r.Group("api/v1")
	//{
	//	// User模块的路由接口
	//	router.PUT("user/:id", v1.UpdateUser)
	//	router.DELETE("user/:id", v1.DeleteUser)
	//
	//	router.POST("user/add", v1.AddUser)
	//	router.GET("users", v1.GetUsers)
	//}
	r.Run(utils.HttpPort)
}
