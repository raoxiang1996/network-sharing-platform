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

		//Upload
		auth.POST("upload", v1.Upload)

		//Course
		auth.POST("course", v1.AddCourse)
		auth.DELETE("course/:id", v1.DeleteCourse)
		auth.PUT("course/:id", v1.EditCourse)

		//Lesson
		auth.POST("lesson", v1.AddLesson)
		auth.DELETE("lesson", v1.DeleteLesson)
		auth.PUT("lesson/:id", v1.EditLesson)
	}

	router := r.Group("api/v1")
	{
		//User
		router.POST("user", v1.AddUser)

		//Login
		router.POST("login/admin", v1.Login)
		router.POST("login/front", v1.FrontLogin)

		//course
		router.GET("course/:id", v1.GetCourse)
		router.GET("courses", v1.GetCourses)
	}
	r.Run(utils.HttpPort)
}
