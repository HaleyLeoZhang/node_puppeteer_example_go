package comic

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (d *Dao) GetImageList(pageId int) (*[]model.ComicImage, error) {
	imageList := make([]model.ComicImage, 0)
	maps := make(map[string]interface{})
	maps["page_id"] = pageId
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err := d.db.Where(maps).Find(&imageList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return &imageList, err
	}

	return &imageList, nil
}
