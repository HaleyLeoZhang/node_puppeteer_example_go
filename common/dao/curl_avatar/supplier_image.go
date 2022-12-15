package curl_avatar

import (
	"context"
	dbTool "github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/po"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (d *Dao) SupplierImageListWithFields(ctx context.Context, chapterId int, fields string) (list []*po.SupplierImage, err error) {
	cond := &dbTool.DBConditions{
		Select: fields,
		And: map[string]interface{}{
			"status = ?":     constant.BASE_TABLE_ONLINE,
			"related_id = ?": chapterId,
		},
		Order: "sequence ASC",
	}
	return d.SupplierImageListByCondition(ctx, cond)
}

func (d *Dao) SupplierImageListByCondition(ctx context.Context, conditions *dbTool.DBConditions) (list []*po.SupplierImage, err error) {
	err = dbTool.Context(ctx, d.db)
	if err != nil {
		return
	}

	var (
		res   = &po.SupplierImage{}
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

func (d *Dao) SupplierImageGetOne(ctx context.Context, id int) (res *po.SupplierImage, err error) {
	res = &po.SupplierImage{}
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
