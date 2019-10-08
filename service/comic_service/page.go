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
	Channel  int
	SourceID int
	ID       int
}

func (p *PageParam) GetList() ([]*models.ComicPages, error) {
	var (
		PageList []*models.ComicPages
		err      error
	)

	cache := caches.ComicPage{
		Channel:  p.Channel,
		SourceID: p.SourceID,
	}
	PageList, err = cache.Get()
	if nil != PageList {
		return PageList, nil
	}

	PageList, err = models.GetPageList(p.Channel, p.SourceID, p.getMaps())
	if err != nil {
		return nil, err
	}

	cache.Save(PageList)

	return PageList, nil
}

func (p *PageParam) GetInfo() (*models.ComicPages, error) {
	var (
		PageInfo *models.ComicPages
		err      error
	)

	PageInfo, err = models.GetPageInfo(p.ID, p.getMaps())
	if err != nil {
		return nil, err
	}

	return PageInfo, nil
}

func (p *PageParam) GetNextInfo() (*models.ComicPages, error) {
	var (
		NextPageInfo *models.ComicPages
		err          error
	)

	NextPageInfo, err = models.GetNextPageInfo(p.Channel, p.SourceID, p.ID, p.getMaps())
	if err != nil {
		return nil, err
	}

	return NextPageInfo, nil
}

func (p *PageParam) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO

	return maps
}
