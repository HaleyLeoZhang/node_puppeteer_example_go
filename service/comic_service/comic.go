package comic_service

// ----------------------------------------------------------------------
// 漫画列表-服务层
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"node_puppeteer_example_go/caches"
	"node_puppeteer_example_go/models"
	"node_puppeteer_example_go/pkg/e"
)

type ComicParam struct {
	PageNum  int
	PageSize int
	Channel  int
	SourceID int
}

const (
	isOnlineYes = 1
	isOnlineNo  = 0
)

func (c *ComicParam) GetList() ([]*models.Comics, error) {
	var (
		ComicList []*models.Comics
	)

	ComicList, err := models.GetComicList(c.PageNum, c.PageSize, c.getMaps())
	if err != nil {
		return nil, err
	}
	return ComicList, nil
}

func (c *ComicParam) Count() (int, error) {
	return models.GetComicTotal(c.getMaps())
}

func (c *ComicParam) GetInfo() (*models.Comics, error) {
	var (
		OneComic *models.Comics
		err      error
	)
	// 缓存
	cache := caches.ComicInfo{
		Channel:  c.Channel,
		SourceID: c.SourceID,
	}
	OneComic, err = cache.Get()
	if nil != OneComic {
		return OneComic, nil
	}
	// 模型
	OneComic, err = models.GetComicInfo(c.Channel, c.SourceID, c.getMaps())
	if err != nil {
		return nil, err
	}

	cache.Save(OneComic)

	return OneComic, nil
}

func (c *ComicParam) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO
	maps["is_online"] = isOnlineYes

	return maps
}
