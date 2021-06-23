package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"University-Information-Website/middleware"
	"University-Information-Website/model"
	"University-Information-Website/utils/errmsg"
)

//超级管理员登陆
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(&data)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	token, code := middleware.SetToken(data.ID, data.Username)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

//学生及老师登陆
func FrontLogin(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckFrontLogin(&data)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	token, code := middleware.SetToken(data.ID, data.Username)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"role":    model.GetAuthority(data.ID),
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
