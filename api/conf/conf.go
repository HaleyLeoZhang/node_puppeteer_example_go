package conf

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"node_puppeteer_example_go/component/driver/db"
	"node_puppeteer_example_go/component/driver/httpserver"
	"node_puppeteer_example_go/component/driver/owngin"
	"node_puppeteer_example_go/component/driver/ownlog"
	"node_puppeteer_example_go/component/driver/redis"
)

var (
	Conf     = &Config{}
	confPath string
)

// Config struct
type Config struct {
	ServiceName string             `yaml:"serviceName" json:"serviceName"`
	HttpServer  *httpserver.Config `yaml:"httpServer" json:"httpServer"`
	Gin         *owngin.Config     `yaml:"gin" json:"gin"`
	DB          *db.Config         `yaml:"db" json:"db"`
	Redis       *redis.Config      `yaml:"redis" json:"redis"`
	Log         *ownlog.Config     `yaml:"log" json:"log"`
}

func init() {
	flag.StringVar(&confPath, "conf", "", "conf values")
}

func Init() (err error) {
	var yamlFile string
	if confPath != "" {
		yamlFile, err = filepath.Abs(confPath)
	} else {
		yamlFile, err = filepath.Abs("../build/app.yaml")
	}
	if err != nil {
		return
	}
	yamlRead, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlRead, Conf)
	if err != nil {
		return
	}
	go load()
	return
}

func load() {
	// 动态加载配置
}
