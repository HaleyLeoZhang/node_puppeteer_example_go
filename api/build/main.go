package main

import (
	"flag"
	"github.com/HaleyLeoZhang/go-component/driver/bootstrap"
	"github.com/HaleyLeoZhang/go-component/driver/httpserver"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/api/http"
	"node_puppeteer_example_go/api/service"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	oneService := service.New(conf.Conf)
	xlog.Init(conf.Conf.Log)

	app := bootstrap.New()
	app.Start(func() { // 此部分代码，请勿阻塞进程
		// 通知错误
		//err = nil
		//if err != nil {
		//	app.NotifyError <- err
		//}
		gin := xgin.New(conf.Conf.Gin)
		go httpserver.Run(conf.Conf.HttpServer, http.Init(gin, oneService)) // 已配置 recovery 不用处理 panic
		return
	}).Stop(func() {
		oneService.Close()
	})

}
