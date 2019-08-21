package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowVideoService 视频投稿的服务
type ShowVideoService struct {
}

//Showw 创建视频
func (c *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video

	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
