package upload

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"University-Information-Website/utils"
	"University-Information-Website/utils/errmsg"
)

func Upload(file multipart.File, fileSize int64, coursesId string, lessonId string) (string, int) {
	keyToOverwrite := coursesId + lessonId
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", utils.Bucket, keyToOverwrite),
	}

	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.Put(context.Background(), &ret, upToken, keyToOverwrite, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println("err,", err)
		return "", errmsg.ERROR
	}
	return utils.QnSever + ret.Key, errmsg.SUCCESS
}
