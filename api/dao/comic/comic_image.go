package comic

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (d *Dao) GetImageList(ctx context.Context, pageId int) (imageList []*model.ComicImage, err error) {
	imageList = make([]*model.ComicImage, 0)
	imageInfo := &model.ComicImage{}
	maps := make(map[string]interface{})
	maps["page_id"] = pageId
	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

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
