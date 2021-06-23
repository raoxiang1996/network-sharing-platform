package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	upload "University-Information-Website/upload"
	"University-Information-Website/utils/errmsg"
)

func UploadImg(c *gin.Context) {
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
