package comic_service

// ----------------------------------------------------------------------
// 漫画实体服务
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"node_puppeteer_example_go/caches"
	"node_puppeteer_example_go/models"
	"node_puppeteer_example_go/pkg/e"
)

type Comic struct {
	PageNum  int
	PageSize int
	Channel  int
	SourceID int
}

const (
	isOnlineYes = 1
	isOnlineNo  = 0
)

func (c *Comic) GetList() ([]*models.Comics, error) {
	var (
		ComicList []*models.Comics
	)

	ComicList, err := models.GetComicList(c.PageNum, c.PageSize, c.getMaps())
	if err != nil {
		return nil, err
	}
	return ComicList, nil
}

func (c *Comic) Count() (int, error) {
	return models.GetComicTotal(c.getMaps())
}

func (c *Comic) GetInfo() (*models.Comics, error) {
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

func (c *Comic) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO
	maps["is_online"] = isOnlineYes

	return maps
}
