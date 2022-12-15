package curl_avatar

import (
	"context"
	dbTool "github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/po"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (d *Dao) SupplierChapterListWithFields(ctx context.Context, supplierId int, fields string) (list []*po.SupplierChapter, err error) {
	cond := &dbTool.DBConditions{
		Select: fields,
		And: map[string]interface{}{
			"status = ?":     constant.BASE_TABLE_ONLINE,
			"related_id = ?": supplierId,
		},
		Order: "sequence ASC",
	}
	return d.SupplierChapterListByCondition(ctx, cond)
}

func (d *Dao) SupplierChapterListByCondition(ctx context.Context, conditions *dbTool.DBConditions) (list []*po.SupplierChapter, err error) {
	err = dbTool.Context(ctx, d.db)
	if err != nil {
		return
	}

	var (
		res   = &po.SupplierChapter{}
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
