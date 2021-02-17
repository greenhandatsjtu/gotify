package models

//推送消息的格式
type NewMsg struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}
