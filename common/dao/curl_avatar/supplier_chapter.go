package curl_avatar

import (
	"context"
	dbTool "github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"node_puppeteer_example_go/common/constant"
	"node_puppeteer_example_go/common/model/po"
)

func (d *Dao) SupplierChapterList(ctx context.Context, where map[string]interface{}, attr map[string]interface{}) (res []*po.SupplierChapter, err error) {
	res = make([]*po.SupplierChapter, 0)
	comicInfo := &po.SupplierChapter{}

	err = dbTool.Context(ctx, d.db)
	if err != nil {
		return
	}
	chain := d.db

	if v, exist := attr["limit"]; exist {
		chain = chain.Limit(v)
	}
	if v, exist := attr["offset"]; exist {
		chain = chain.Offset(v)
	}
	if v, exist := attr["order_by"]; exist {
		chain = chain.Order(v)
	}
	if v, exist := attr["select"]; exist {
		chain = chain.Select(v)
	}

	err = chain.Table(comicInfo.TableName()).Where(where).Find(&res).Error

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

func (d *Dao) SupplierChapterGetOne(ctx context.Context, id int) (res *po.SupplierChapter, err error) {
	res = &po.SupplierChapter{}
	err = nil

	chain := d.db
	err = chain.Table(res.TableName()).Where("id = ?", id).First(&res).Error
	if gorm.IsRecordNotFoundError(err) {
		res = nil
		return
	}
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (d *Dao) SupplierChapterGetNextOne(ctx context.Context, sequence int, supplierId int) (res *po.SupplierChapter, err error) {
	res = &po.SupplierChapter{}
	err = nil

	chain := d.db
	err = chain.Table(res.TableName()).Where("sequence > ? AND related_id = ? AND status = ? ", sequence, supplierId, constant.BASE_TABLE_ONLINE).
		Order("sequence asc").First(&res).Error
	if gorm.IsRecordNotFoundError(err) {
		res = nil
		return
	}
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
