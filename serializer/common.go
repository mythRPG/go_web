package serializer

// Response 团队基础序列化器
type Response struct {
	Status int         `json:"status"` //状态码
	Data   interface{} `json:"data"`   //传输的内容
	Msg    string      `json:"msg"`    //返回的错误消息
	Error  string      `json:"error`   //代码错误消息
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}
