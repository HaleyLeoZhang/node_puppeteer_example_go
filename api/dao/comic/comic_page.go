package comic

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (d *Dao) GetPageList(ctx context.Context, maps map[string]interface{}) (pageList []*model.ComicPage, err error) {
	pageList = make([]*model.ComicPage, 0)
	pageInfo := &model.ComicPage{}
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(pageInfo.TableName()).
		Where(maps).Not("sequence", 0).Order("sequence ASC").Find(&pageList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return
	}

	return
}

func (d *Dao) GetPageInfo(ctx context.Context, id int) (pageInfo *model.ComicPage, err error) {
	pageInfo = &model.ComicPage{}

	maps := make(map[string]interface{})
	maps["id"] = id
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(pageInfo.TableName()).
		Where(maps).First(&pageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return
	}

	return
}

func (d *Dao) GetMextPageInfo(ctx context.Context, id int) (pageInfo *model.ComicPage, err error) {
	pageInfo = &model.ComicPage{}

	maps := make(map[string]interface{})
	maps["id"] = id
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(pageInfo.TableName()).
		Where(maps).First(&pageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return
	}

	return
}

func (d *Dao) GetNextPageInfo(ctx context.Context, Channel int, SourceId int, Sequence int) (pageInfo *model.ComicPage, err error) {
	pageInfo = &model.ComicPage{}

	maps := make(map[string]interface{})
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO
	err = d.db.Table(pageInfo.TableName()).
		Where("channel = ? And source_id = ? And sequence > ?", Channel, SourceId, Sequence).
		Where(maps).
		Order("sequence ASC").
		First(&pageInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return nil, err
	}

	return pageInfo, nil
}
