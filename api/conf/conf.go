package conf

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"node_puppeteer_example_go/components/driver/DB"
	"node_puppeteer_example_go/components/driver/Redis"
)

var (
	Conf     = &Config{}
	confPath string
)

// Config struct
type Config struct {
	Log         *Redis.Conf
	HttpServer  *Gin.Conf
	DB          *DB.Config
	Redis       *redis.Config
	ServiceName string      `yaml:"serviceName" json:"serviceName"`
}

func init() {
	flag.StringVar(&confPath, "conf", "", "conf values")
}

func Init() (err error) {
	var (
		yamlFile string
	)
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
