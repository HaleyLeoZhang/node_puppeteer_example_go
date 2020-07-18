package conf

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"node_puppeteer_example_go/component/driver/db"
	"node_puppeteer_example_go/component/driver/httpserver"
	"node_puppeteer_example_go/component/driver/owngin"
	"node_puppeteer_example_go/component/driver/redis"
)

var (
	Conf     = &Config{}
	confPath string
)

// Config struct
type Config struct {
	ServiceName string             `yaml:"serviceName"`
	HttpServer  *httpserver.Config `yaml:"httpServer"`
	Gin         *owngin.Config     `yaml:"gin"`
	DB          *db.Config         `yaml:"db"`
	Redis       *redis.Config
	//Log         *Log.Config
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
	fmt.Println(Conf)
	go load()
	return
}

func load() {
	// 动态加载配置
}
