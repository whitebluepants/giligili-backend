package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情
type ShowVideoService struct {
}

// Show 视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
