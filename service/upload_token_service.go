package service

import (
	"giligili/serializer"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// UploadTokenService 视频投稿的服务
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

// Create 创建视频
func (service *UploadTokenService) Post() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg: "OSS配置错误1",
			Error: err.Error(),
		}
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil{
		return serializer.Response{
			Status: 50002,
			Msg: "OSS配置错误2",
			Error: err.Error(),
		}
	}

	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ".png"
	//签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil{
		return serializer.Response{
			Status: 50002,
			Msg: "OSS配置错误3",
			Error: err.Error(),
		}
	}

	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil{
		return serializer.Response{
			Status: 50002,
			Msg: "OSS配置错误4",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
