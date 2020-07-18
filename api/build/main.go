package main

import (
	"flag"
	"node_puppeteer_example_go/api/conf"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	//app := bootstrap.New(bootstrap.Config{
	//	Log:         conf.Conf.Log,
	//	Tracer:      conf.Conf.Tracer,
	//	GinServer:   conf.Conf.GinServer,
	//	ServiceName: conf.Conf.ServiceName,
	//})
	//
	////service process
	//srv := service.New(conf.Conf)
	//defer srv.Close()
	//
	//app.Init(
	//	//http process
	//	bootstrap.HTTPService(func(r *gins.Server, m *metrics.Metrics) {
	//		http.Init(r, srv, m)
	//	}),
	//
	//	//micro process
	//	bootstrap.MicroService(func(s *micro.Service) {
	//		//client
	//		handler.SupplierServiceInit(s.Client())
	//
	//		//grcp
	//		if err := pb.RegisterSupplierHandler(s.Server(), new(handler.SupplierHanders)); err != nil {
	//			xlog.Error(err)
	//		}
	//	}),
	//)
	//
	//if err := app.Run(); err != nil {
	//	panic(err)
	//}
}
