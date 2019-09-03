package comic_service

// ----------------------------------------------------------------------
// 漫画列表-服务层
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/models"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/e"
)

type ComicParam struct {
	PageNum  int
	PageSize int
}

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

func (c *ComicParam) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO

	return maps
}
