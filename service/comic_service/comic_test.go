package comic_service

import (
	"fmt"
	"node_puppeteer_example_go/models"
	"node_puppeteer_example_go/pkg/gredis"
	"node_puppeteer_example_go/pkg/logging"
	"node_puppeteer_example_go/pkg/setting"
	"testing"
)

func TestMain(m *testing.M) {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	m.Run()
}

func TestCore(t *testing.T) {
	t.Run("comicGetInfo", comicGetInfo)
}

func comicGetInfo(t *testing.T) {
	comic := Comic{
		Channel:  3,
		SourceID: 10319,
	}
	list, err := comic.GetInfo()
	if nil == err {
		logging.Info("list结果...", fmt.Sprintf("%v", list))
	} else {
		t.Fatalf("数据不存在: %v", err)
	}
}
