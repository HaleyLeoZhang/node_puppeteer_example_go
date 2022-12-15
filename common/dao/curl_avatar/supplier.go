package curl_avatar

import (
	"context"
	dbTool "github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/po"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (d *Dao) SupplierListForIndexWithFields(ctx context.Context, supplierIds []int, fields string) (list []*po.Supplier, err error) {
	cond := &dbTool.DBConditions{
		Select: fields,
		And: map[string]interface{}{
			"status = ?": constant.BASE_TABLE_ONLINE,
			"id in (?)":  supplierIds,
		},
	}
	return d.SupplierListByCondition(ctx, cond)
}

func (d *Dao) SupplierOneForChapterWithFields(ctx context.Context, comicId int, fields string) (item *po.Supplier, err error) {
	cond := &dbTool.DBConditions{
		Select: fields,
		And: map[string]interface{}{
			"status = ?":     constant.BASE_TABLE_ONLINE,
			"related_id = ?": comicId,
		},
		Limit: 1,
		Order: "weight DESC,id DESC",
	}
	list, err := d.SupplierListByCondition(ctx, cond)
	if err != nil {
		return
	}
	if len(list) > 0 {
		item = list[0]
		return
	}
	return
}

func (d *Dao) SupplierListByCondition(ctx context.Context, conditions *dbTool.DBConditions) (list []*po.Supplier, err error) {
	err = dbTool.Context(ctx, d.db)
	if err != nil {
		return
	}

	var (
		res   = &po.Supplier{}
		chain = d.db
	)
	chain = chain.Table(res.TableName())
	chain = conditions.Fill(chain)
	err = chain.Find(&list).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
		return
	}
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (d *Dao) SupplierGetOne(ctx context.Context, id int) (res *po.Supplier, err error) {
	res = &po.Supplier{}
	err = nil

	chain := d.db
	err = chain.Table(res.TableName()).Where("id = ?", id).First(&res).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
		return
	}
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
