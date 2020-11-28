package comic

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	constant2 "node_puppeteer_example_go/common/constant"
	"node_puppeteer_example_go/common/model/po"
)

func (d *Dao) GetPageList(ctx context.Context, maps map[string]interface{}) (pageList []*po.ComicPage, err error) {
	pageList = make([]*po.ComicPage, 0)
	pageInfo := &po.ComicPage{}
	maps["is_deleted"] = constant2.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(pageInfo.TableName()).
		Where(maps).Not("sequence", 0).Order("sequence ASC").Find(&pageList).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

func (d *Dao) GetPageInfo(ctx context.Context, id int) (pageInfo *po.ComicPage, err error) {
	pageInfo = &po.ComicPage{}

	maps := make(map[string]interface{})
	maps["id"] = id
	maps["is_deleted"] = constant2.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(pageInfo.TableName()).
		Where(maps).First(&pageInfo).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

func (d *Dao) GetNextPageInfo(ctx context.Context, Channel int, SourceId int, Sequence int) (pageInfo *po.ComicPage, err error) {
	pageInfo = &po.ComicPage{}

	maps := make(map[string]interface{})
	maps["is_deleted"] = constant2.TABLE_BASE_IS_DELETED_NO
	err = d.db.Table(pageInfo.TableName()).
		Where("channel = ? And source_id = ? And sequence > ?", Channel, SourceId, Sequence).
		Where(maps).
		Order("sequence ASC").
		First(&pageInfo).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}

	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return pageInfo, nil
}
