package service

import (
	"singo/model"
	"singo/serializer"
)

// CreatVideoService 视频投稿的服务
type CreatVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=8,max=40"`
}

//Creat 创建视频
func (c *CreatVideoService) Create() serializer.Response {
	video := model.Video{
		Title: c.Title,
		Info:  c.Info,
	}
	err := model.DB.Create(&video).Error
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
