package service

import (
	"singo/model"
	"singo/serializer"
)

//  视频更新的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=8,max=40"`
}

//Showw 更新视频
func (s *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video

	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	video.Title = s.Title
	video.Info = s.Info

	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频修改失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
