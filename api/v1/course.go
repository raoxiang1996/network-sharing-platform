package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"University-Information-Website/model"
	"University-Information-Website/utils/errmsg"
)

func AddCourse(c *gin.Context) {
	var data model.Courses
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	code := model.CreateCourse(&data, data.ID)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteCourse(c *gin.Context) {
	var data model.Courses
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

}

func AddLesson(c *gin.Context) {
	var data model.Lesson
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	code := model.InsertLesson(&data, data.CoursesId)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
