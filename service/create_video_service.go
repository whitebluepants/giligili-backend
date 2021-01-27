package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
	Url   string `form:"url" json:"url"`
	Avatar string `form:"avatar" json:"avatar"`
}

// Create 创建视频
func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
		Url:   service.Url,
		Avatar: service.Avatar,
		UserID: user.ID,
	}

	//fmt.Printf("%+v\n", video)
	// DB的Create函数会修改video的值, 因为往里面传的是指针
	err := model.DB.Create(&video).Error
	//fmt.Printf("%+v\n", video)
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
