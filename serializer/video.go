package serializer

import "singo/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`         //自动编号
	Title     string `json:"title"`      //标签
	Info      string `json:"info"`       //信息
	CreatedAt int64  `json:"created_at"` //时间戳
}

// BuildVideo 序列化用户
func BuildVideo(v model.Video) Video {
	return Video{
		ID:        v.ID,
		Title:     v.Title,
		Info:      v.Info,
		CreatedAt: v.CreatedAt.Unix(),
	}
}

func BuildVideos(v []model.Video) (videos []Video) {

	for _, v := range v {
		video := BuildVideo(v)
		videos = append(videos, video)
	}

	return videos
}
