package models

// ----------------------------------------------------------------------
// 漫画列表-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/logging"
	"github.com/jinzhu/gorm"

	"encoding/json"
)

type Comics struct {
	ID        string `json:"id"`
	Channel   string `json:"channel"`
	ComicID   string `json:"comic_id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	Intro     string `json:"intro"`
	IsDeleted string `json:"is_deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func GetComicList(pageNum int, pageSize int, maps interface{}) ([]*Comics, error) {
	var ComicList []*Comics
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&ComicList).Error

	page_data, err := json.Marshal(ComicList)

	logging.Info("maps ", maps)
	logging.Info("pageNum ", pageNum)
	logging.Info("pageSize ", pageSize)
	logging.Info("page_data ", string(page_data))

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ComicList, nil
}

func GetComicTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Comics{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
