package utils

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func InitApp() {
	home, _ := homedir.Dir()
	dir := filepath.Join(home, ".gotify")
	_ = os.Mkdir(dir, os.ModePerm)
	config, _ := ioutil.ReadFile(filepath.Join("config", "config.json"))
	if err := ioutil.WriteFile(filepath.Join(dir, "config.json"), config, 0644); err != nil {
		log.Fatal("init config error: ", err)
	}
	imagePath := filepath.Join(dir, "images")
	_ = os.Mkdir(imagePath, os.ModePerm)
	image, _ := ioutil.ReadFile(filepath.Join("images", "logo.png"))
	if err := ioutil.WriteFile(filepath.Join(imagePath, "logo.png"), image, 0644); err != nil {
		log.Fatal(err)
	}
	pluginPath := filepath.Join(dir, "plugins")
	_ = os.Mkdir(pluginPath, os.ModePerm)
	files, _ := ioutil.ReadDir("plugins")
	for _, f := range files {
		plugin, _ := ioutil.ReadFile(filepath.Join("plugins", f.Name()))
		if err := ioutil.WriteFile(filepath.Join(pluginPath, f.Name()), plugin, 0744); err != nil {
			log.Fatal(err)
		}
	}
	//db, _ := ioutil.ReadFile("gotify.db")
	//if err := ioutil.WriteFile(filepath.Join(dir, "gotify.db"), db, 0644); err != nil {
	//	log.Fatal("create db error:", err)
	//}
}
