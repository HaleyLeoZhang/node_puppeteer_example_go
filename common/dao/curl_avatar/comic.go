package curl_avatar

import (
	"context"
	dbTool "github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/po"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (d *Dao) ComicListForIndexWithFields(ctx context.Context, limit, offset int, orderBy, fields string) (list []*po.Comic, err error) {
	cond := &dbTool.DBConditions{
		Select: fields,
		And: map[string]interface{}{
			"status = ?": constant.BASE_TABLE_ONLINE,
		},
		Limit:  limit,
		Offset: offset,
		Order:  orderBy,
	}
	return d.comicListByCondition(ctx, cond)
}

func (d *Dao) comicListByCondition(ctx context.Context, conditions *dbTool.DBConditions) (list []*po.Comic, err error) {
	err = dbTool.Context(ctx, d.db)
	if err != nil {
		return
	}

	var (
		res   = &po.Comic{}
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

func (d *Dao) ComicGetOne(ctx context.Context, id int) (res *po.Comic, err error) {
	res = &po.Comic{}
	err = nil

	err = dbTool.Context(ctx, d.db)
	if err != nil {
		return
	}

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
