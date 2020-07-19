package comic

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (d *Dao) GetPageList(ctx context.Context, maps map[string]interface{}) (*[]model.ComicPage, error) {
	pageList := make([]model.ComicPage, 0)
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err := d.db.Where(maps).Not("sequence", 0).Order("sequence ASC").Find(&pageList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return &pageList, err
	}

	return &pageList, nil
}

func (d *Dao) GetPageInfo(ctx context.Context, id int) (*model.ComicPage, error) {
	pageInfo := &model.ComicPage{}

	maps := make(map[string]interface{})
	maps["id"] = id
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err := d.db.Where(maps).First(&pageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return nil, err
	}

	return pageInfo, nil
}

func (d *Dao) GetMextPageInfo(ctx context.Context, id int) (*model.ComicPage, error) {
	pageInfo := &model.ComicPage{}

	maps := make(map[string]interface{})
	maps["id"] = id
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err := d.db.Where(maps).First(&pageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return pageInfo, err
	}

	return pageInfo, nil
}

func (d *Dao) GetNextPageInfo(ctx context.Context, Channel int, SourceId int, Sequence int) (*model.ComicPage, error) {
	pageInfo := &model.ComicPage{}

	maps := make(map[string]interface{})
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO
	err := d.db.Where("channel = ? And source_id = ? And sequence > ?", Channel, SourceId, Sequence).
		Where(maps).
		Order("sequence ASC").
		First(&pageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return nil, err
	}

	return pageInfo, nil
}
