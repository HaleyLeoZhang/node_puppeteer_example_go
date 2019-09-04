package comic_service

// ----------------------------------------------------------------------
// 漫画列表-服务层
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	// "encoding/json"

	"github.com/HaleyLeoZhang/node_puppeteer_example_go/caches"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/models"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/e"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/gredis"
	// "github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/logging"
)

type PageParam struct {
	Channel int
	ComicID int
}

func (p *PageParam) GetList() ([]*models.Pages, error) {
	var (
		PageList []*models.Pages
		err      error
	)

	cache := caches.ComicPage{
		Channel: p.Channel,
		ComicID: p.ComicID,
	}
	PageList, err = cache.Get()
	if nil != PageList {
		return PageList, nil
	}

	PageList, err = models.GetPageList(p.Channel, p.ComicID, p.getMaps())
	if err != nil {
		return nil, err
	}

	cache.Save(PageList)

	return PageList, nil
}

func (p *PageParam) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO

	return maps
}
