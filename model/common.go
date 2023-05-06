package model

// RemoveRequest 前端需要实现批量删除的请求结构体
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
