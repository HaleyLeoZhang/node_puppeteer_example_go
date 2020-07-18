package models

// ----------------------------------------------------------------------
// 漫画章节对应图片列表-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"github.com/jinzhu/gorm"
)

type ComicImages struct {
	ID       string `json:"id"`
	PageID   string `json:"page_id"`
	Sequence string `json:"sequence"`
	Src      string `json:"src"`
	Progress string `json:"progress"`
	// IsDeleted string `json:"is_deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func GetImageList(PageID int, maps interface{}) ([]*ComicImages, error) {
	var ImageList []*ComicImages

	query_maps := make(map[string]interface{})
	query_maps["page_id"] = PageID

	err := db.Where(query_maps).Where(maps).Order("sequence asc").Find(&ImageList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ImageList, nil
}
