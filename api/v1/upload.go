package v1

import (
	upload "University-Information-Website/Upload"
	"University-Information-Website/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("err:", err)
		code := errmsg.ERROR
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	fileSize := fileHeader.Size
	url, code := upload.Upload(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"url":    url,
	})
}
