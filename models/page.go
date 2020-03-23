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

type ComicPages struct {
	ID       int    `json:"id"`
	Channel  int    `json:"channel"`
	SourceID int    `json:"source_id"`
	Sequence int    `json:"sequence"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	Progress int    `json:"progress"`
	// IsDeleted int `json:"is_deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func GetPageList(Channel int, SourceID int, maps interface{}) ([]*ComicPages, error) {
	var PageList []*ComicPages

	query_maps := make(map[string]interface{})
	query_maps["channel"] = Channel
	query_maps["source_id"] = SourceID

	err := db.Where(query_maps).Where(maps).Not("sequence", 0).Order("sequence asc").Find(&PageList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return PageList, nil
}

func GetPageInfo(ID int, maps interface{}) (*ComicPages, error) {
	var PageInfo ComicPages

	query_maps := make(map[string]interface{})
	query_maps["id"] = ID

	err := db.Where(query_maps).Where(maps).First(&PageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &PageInfo, nil
}

func GetNextPageInfo(Channel int, SourceID int, ID int, maps interface{}) (*ComicPages, error) {
	var PageInfo ComicPages

	err := db.Where("channel = ? And source_id = ? And id > ?", Channel, SourceID, ID).Where(maps).Order("sequence desc").First(&PageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &PageInfo, nil
}
