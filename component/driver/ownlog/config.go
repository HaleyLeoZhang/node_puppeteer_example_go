package ownlog

type Config struct {
	Name   string `yaml:"name" json:"name"`     // 日志服务名  写入文件或者发到日志收集服务时使用
	Stdout bool   `yaml:"stdout" json:"stdout"` // 是否输出
	Dir    string `yaml:"dir" json:"dir"`       // 如果不输出，存储位置
}
