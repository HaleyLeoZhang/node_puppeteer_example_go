package models

// ----------------------------------------------------------------------
// 漫画章节列表-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
import (
	"github.com/jinzhu/gorm"
)

type Pages struct {
	ID       int    `json:"id"`
	Channel  int    `json:"channel"`
	ComicID  int    `json:"comic_id"`
	Sequence int    `json:"sequence"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	Progress int    `json:"progress"`
	// IsDeleted int `json:"is_deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func GetPageList(Channel int, ComicID int, maps interface{}) ([]*Pages, error) {
	var PageList []*Pages

	query_maps := make(map[string]interface{})
	query_maps["channel"] = Channel
	query_maps["comic_id"] = ComicID

	err := db.Where(query_maps).Where(maps).Not("sequence", 0).Order("sequence asc").Find(&PageList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return PageList, nil
}
