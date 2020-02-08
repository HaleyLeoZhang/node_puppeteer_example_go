package comic_service

// ----------------------------------------------------------------------
// 漫画列表-服务层
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"node_puppeteer_example_go/models"
	"node_puppeteer_example_go/pkg/e"
)

type ImageParam struct {
	PageID int
}

func (i *ImageParam) GetList() ([]*models.ComicImages, error) {
	var (
		ImageList []*models.ComicImages
	)

	ImageList, err := models.GetImageList(i.PageID, i.getMaps())
	if err != nil {
		return nil, err
	}
	return ImageList, nil
}

func (i *ImageParam) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_deleted"] = e.DATA_IS_DELETED_NO

	return maps
}
