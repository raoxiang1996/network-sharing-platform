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

		auth.GET("users", v1.GetUsers)
		auth.PUT("upload/img", v1.UploadImg)
		auth.PUT("upload/video", v1.UploadImg)
	}

	router := r.Group("api/v1")
	{
		//User
		router.POST("user/add", v1.AddUser)

		//Login
		router.POST("login/admin", v1.Login)
		router.POST("login/front", v1.FrontLogin)
	}
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
