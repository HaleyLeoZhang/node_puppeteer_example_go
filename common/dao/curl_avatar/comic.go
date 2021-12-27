package curl_avatar

import (
	"context"
	dbTool "github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/po"
)

func (d *Dao) ComicList(ctx context.Context, where map[string]interface{}, attr map[string]interface{}) (res []*po.Comic, err error) {
	res = make([]*po.Comic, 0)
	comicInfo := &po.Comic{}

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
