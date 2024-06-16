/*
错误数据
*/
package types

// 服务端请求错误返回消息格式
type ErrData struct {
	Event string `json:"event"`
	Code  string `json:"code"`
	Msg   string `json:"msg"`
}
