package bootstrap

import (
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/component/driver/ownlog"
	"os"
	"os/signal"
	"syscall"
)

type Bootstrap struct {
	ExitSignal  chan os.Signal
	NotifyError chan error
}

func New() *Bootstrap {
	return &Bootstrap{}
}

/**
 * 监听开始
 */
func (b *Bootstrap) Start(callback func()) *Bootstrap {
	defer func() {
		if p := recover(); p != nil {
			ownlog.Errorf("App(%v).Panic.(%v)", conf.Conf.ServiceName, p)
		}
	}()
	ownlog.Info("Bootstrap.ing")
	callback()
	ownlog.Info("Bootstrap.done")
	return b
}

/**
 * 监听结束
 */
func (b *Bootstrap) Stop(callback func()) {
	b.ExitSignal = make(chan os.Signal)
	// 监听指定信号 ctrl+c 、kill 进程Pid
	// - 请不要用 kil -9 程序监听不到退出
	// - 调试的时候 kill 目标请杀掉对应 tmp 进程即可
	signal.Notify(b.ExitSignal, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case s := <-b.ExitSignal: // 阻塞直至有信号传入
			ownlog.Infof("Received Exit Signal: %v", s)
			callback()
			os.Exit(0)
		case err := <-b.NotifyError:
			ownlog.Errorf("Bootstrap.Start.Err(%+v)", err)
			b.ExitSignal <- syscall.SIGINT
			//case <-time.After(time.Second * 3): // 检测进程是否存活，暂不需要
			//	ownlog.Info("Bootstrap.Loop.Alive")
		}
		ownlog.Info("Bootstrap.test")
	}

}
