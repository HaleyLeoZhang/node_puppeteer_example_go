package comic_service

// ----------------------------------------------------------------------
// 漫画章节服务
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	// "encoding/json"

	"node_puppeteer_example_go/caches"
	"node_puppeteer_example_go/models"
	"node_puppeteer_example_go/pkg/e"
	// "node_puppeteer_example_go/pkg/gredis"
	// "node_puppeteer_example_go/pkg/logging"
)

type Page struct {
	Channel  int
	SourceID int
	Sequence int
	ID       int
}

func (p *Page) GetList() ([]*models.ComicPages, error) {
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

func (p *Page) GetInfo() (*models.ComicPages, error) {
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

func (p *Page) GetNextInfo() (*models.ComicPages, error) {
	var (
		NextPageInfo *models.ComicPages
		err          error
	)

	NextPageInfo, err = models.GetNextPageInfo(p.Channel, p.SourceID, p.Sequence, p.getMaps())
	if err != nil {
		return nil, err
	}

	return NextPageInfo, nil
}

func (p *Page) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO

	return maps
}
