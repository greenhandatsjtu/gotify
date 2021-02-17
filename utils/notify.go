package utils

import (
	"github.com/gen2brain/beeep"
	"github.com/greenhandatsjtu/gotify/models"
	log "github.com/sirupsen/logrus"
	"time"
)

var err error

//系统通知，从消息队列中取消息弹出系统通知
func Notify(msgs chan models.NewMsg, iconPath string) {
	//阻塞地等待消息队列
	for msg := range msgs {
		log.Info(msg.Title)
		//调用API弹出桌面通知
		err = beeep.Notify(msg.Title, msg.Message, iconPath)
		if err != nil {
			log.Warn(err)
		}

		//byteMsg, _ := json.Marshal(msg)
		//m.Broadcast(byteMsg) //用于websocket向前端发送Msg，已弃用

		//适当休眠
		time.Sleep(time.Millisecond * 2550)
	}
}
