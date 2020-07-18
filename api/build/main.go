package main

import (
	"flag"
	"fmt"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/api/http"
	"node_puppeteer_example_go/component/driver/httpserver"
	"node_puppeteer_example_go/component/driver/owngin"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	fmt.Printf("conf.Conf.Gin %+v", conf.Conf.Gin)
	gin := owngin.New(conf.Conf.Gin)
	httpserver.Run(conf.Conf.HttpServer, http.Init(conf.Conf, gin))
	//app := bootstrap.New(bootstrap.Config{
	//	Log:         conf.Conf.Log,
	//	Tracer:      conf.Conf.Tracer,
	//	GinServer:   conf.Conf.GinServer,
	//	ServiceName: conf.Conf.ServiceName,
	//})
}
