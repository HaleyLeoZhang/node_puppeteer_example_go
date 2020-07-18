package comic

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/driver/db"
)

func (d *Dao) GetComicList(page int, size int, maps map[string]interface{}) (*[]model.Comic, error) {
	comicList := make([]model.Comic, 0)
	offset, size := db.GetPageInfo(page, size)

	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err := d.db.Where(maps).Offset(offset).Order("weight DESC").Limit(size).Find(&comicList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return &comicList, err
	}

	return &comicList, nil
}

func (d *Dao) GetComicInfo(Channel int, SourceID int) (*model.Comic, error) {
	comicInfo := &model.Comic{}

	maps := make(map[string]interface{})
	maps["channel"] = Channel
	maps["source_id"] = SourceID

	err := d.db.Where(maps).First(&comicInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return comicInfo, nil
}
