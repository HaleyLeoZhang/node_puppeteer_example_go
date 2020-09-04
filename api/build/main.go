package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/api/http"
	"node_puppeteer_example_go/api/service"
	"node_puppeteer_example_go/component/driver/bootstrap"
	"node_puppeteer_example_go/component/driver/httpserver"
	"node_puppeteer_example_go/component/driver/owngin"
	"node_puppeteer_example_go/component/driver/ownlog"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	confString, err := json.Marshal(conf.Conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nconf.Conf %+v\n\n", string(confString))

	oneService := service.New(conf.Conf)
	ownlog.Init(conf.Conf.Log)

	app := bootstrap.New()
	app.Start(func() { // 此部分代码，请勿阻塞进程
		// 通知错误
		//err = nil
		//if err != nil {
		//	app.NotifyError <- err
		//}
		gin := owngin.New(conf.Conf.Gin)
		go httpserver.Run(conf.Conf.HttpServer, http.Init(gin, oneService)) // 已配置 recovery 不用处理 panic
		return
	}).Stop(func() {
		oneService.Close()
	})

}
