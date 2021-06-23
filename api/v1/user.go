package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"University-Information-Website/middleware"
	"University-Information-Website/model"
	"University-Information-Website/utils/errmsg"
)

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	code := model.CheckUser(&data)
	if code == errmsg.SUCCESS {
		model.InsertUser(&data)
	} else {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var code int
	id := c.Param("id")
	tokenHeader := c.Request.Header.Get("Authorization")
	userId, _, role, code := middleware.ParseToken(tokenHeader)

	if code != errmsg.SUCCESS {
		code = errmsg.ERROR_USER_DEL_ERROR
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	if role >= 2 && userId == id {
		code = model.DeleteUser(id)
	} else if role < 2 && userId != id {
		code = model.DeleteUser(id)
	} else {
		code = errmsg.ERROR_USER_DEL_ERROR
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetUsers(pageSize, pageNum)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改用户
func UpdateUser(c *gin.Context) {
	var data model.User
	id := c.Param("_id")
	c.ShouldBindJSON(&data)
	code := model.CheckUser(&data)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	model.UpdateUser(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetAuthority(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"role": model.GetAuthority(id),
	})
}
