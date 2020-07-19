package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Name           string        `yaml:"name"` // 用于 Trace 识别
	Ip             string        `yaml:"ip"`
	Port           int           `yaml:"port"`
	Timeout        string        `yaml:"timeout"`
	ReadTimeout    time.Duration `yaml:"readTimeout"`
	WriteTimeout   time.Duration `yaml:"writeTimeout"`
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"`
}

func Run(c *Config, routersInit *gin.Engine) {
	addrString := fmt.Sprintf("%s:%v", c.Ip, c.Port)

	server := &http.Server{
		Addr:           addrString,
		Handler:        routersInit,
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		MaxHeaderBytes: c.MaxHeaderBytes,
	}
	log.Printf("Start http server listening %s", addrString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("HttpServer.Err %+v", err)
	}
}
