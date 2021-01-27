package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情
type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 创建视频
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	total := 0

	if service.Limit == 0{
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil{
		return serializer.Response{
			Status: 50000,
			Msg: "数据库连接错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil{
		return serializer.Response{
			Status: 50000,
			Msg: "数据库连接错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
