package sync

import (
	"github.com/greenhandatsjtu/notifier/models"
	"github.com/greenhandatsjtu/notifier/utils"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"path/filepath"
	"time"
)

//每当计时器发出信号就循环执行插件命令
func LoopSyncPlugin(plugin models.Plugin, msgs chan models.NewMsg) {
	timer := make(chan bool)
	defer close(timer)
	//设置定时器
	go utils.TickTock(timer, time.Duration(plugin.Frequency)*time.Minute)

	for {
		<-timer
		log.Infof("执行插件%s命令: %s\n", plugin.Name, plugin.Exec)
		//执行命令
		home, _ := homedir.Dir()
		cmd := exec.Command(filepath.Join(home, ".gotify", "plugins", plugin.Exec))
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Warnf("执行出错: %s\n", err)
		}
		if "" != string(output) {
			msgs <- models.NewMsg{
				Title:   plugin.Name,
				Message: string(output),
			}
		}
	}
}
