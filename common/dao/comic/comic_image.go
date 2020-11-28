package comic

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	constant2 "node_puppeteer_example_go/common/constant"
	"node_puppeteer_example_go/common/model/po"
)

func (d *Dao) GetImageList(ctx context.Context, pageId int) (imageList []*po.ComicImage, err error) {
	imageList = make([]*po.ComicImage, 0)
	imageInfo := &po.ComicImage{}
	maps := make(map[string]interface{})
	maps["page_id"] = pageId
	maps["is_deleted"] = constant2.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(imageInfo.TableName()). // 以此减小反射的开销
							Where(maps).Find(&imageList).Error

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
