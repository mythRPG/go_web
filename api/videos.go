package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

//视频投稿
func CreatVideo(c *gin.Context) {
	var service service.CreatVideoService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse((err)))
	}
}

//视频详情接口
func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse((err)))
	}
}

//视频列表接口
func ListVideo(c *gin.Context) {
	var service service.ListVideoService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse((err)))
	}
}

//视频更新接口
func UpdateVideo(c *gin.Context) {
	var service service.UpdateVideoService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse((err)))
	}
}

func DeleteVideo(c *gin.Context) {
	var service service.DeleteVideoService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse((err)))
	}
}
