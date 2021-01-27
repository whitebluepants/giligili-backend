package serializer

// Response 团队基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total uint		  `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

func BuildListResponse(items interface{}, total uint) Response{
	return Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}