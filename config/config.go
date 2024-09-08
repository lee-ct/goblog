package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type TomlConfig struct {
	Viewer       Viewer
	SystemConfig SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *TomlConfig

func init() {
	Cfg = new(TomlConfig)
	var err error
	Cfg.SystemConfig.CurrentDir, err = os.Getwd()
	Cfg.SystemConfig.AppName = "lct'blog"
	Cfg.SystemConfig.Version = 1.0
	if err != nil {
		log.Println(err)
	}
	_, err = toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		log.Println("配置文件解析错误", err)
	}
}
