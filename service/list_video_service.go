package service

import (
	"singo/model"
	"singo/serializer"
)

//  视频列表的服务
type ListVideoService struct {
}

//Showw 视频列表
func (l *ListVideoService) List() serializer.Response {
	var videos []model.Video

	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
