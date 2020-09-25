package comic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
)

func (d *Dao) GetComicList(ctx context.Context, page int, size int, maps map[string]interface{}) (comicList []*model.Comic, err error) {
	comicList = make([]*model.Comic, 0)
	comicInfo := &model.Comic{}

	offset, size := db.GetPageInfo(page, size)

	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err = d.db.Table(comicInfo.TableName()).
		Where(maps).Offset(offset).Order("weight DESC").Limit(size).Find(&comicList).Error

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

func (d *Dao) GetComicInfo(ctx context.Context, Channel int, SourceID int) (comicInfo *model.Comic, err error) {
	comicInfo = &model.Comic{}

	maps := make(map[string]interface{})
	maps["channel"] = Channel
	maps["source_id"] = SourceID

	err = d.db.Table(comicInfo.TableName()).
		Where(maps).First(&comicInfo).Error

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

// --------------------------------
// 漫画基础信息缓存
// --------------------------------
func cacheKeyGetComicInfo(Channel int, SourceID int) string {
	prefix := constant.CACHE_PREFIX_TYPE_FRONT
	namespace := fmt.Sprintf("page_detail_comic_%v_%v", Channel, SourceID)
	version := "v1"
	suffix := constant.CACHE_SUFFIX_TYPE_CACHED
	cacheKey := fmt.Sprintf("%v:%v:%v:%v", prefix, namespace, version, suffix)
	return cacheKey
}

func (d *Dao) CacheSetComicInfo(ctx context.Context, comic *model.Comic) error {
	cacheKey := cacheKeyGetComicInfo(comic.Channel, comic.SourceID)
	byteValue, err := json.Marshal(comic)
	if nil != err {
		err = errors.WithStack(err)
		return err
	}
	ttl := 300 // 缓存时间，单位，秒

	conn := d.redis.Get()
	defer conn.Close()
	_, err = conn.Do("SET", cacheKey, byteValue, "EX", ttl)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return err
}

func (d *Dao) CacheGetComicInfo(ctx context.Context, Channel int, SourceID int) (*model.Comic, error) {
	cacheKey := cacheKeyGetComicInfo(Channel, SourceID)
	conn := d.redis.Get()
	defer conn.Close()
	byteValue, err := conn.Do("GET", cacheKey)
	if err != nil {
		return nil, err
	}
	comic := &model.Comic{}
	if nil == byteValue {
		return nil, nil
	}
	json.Unmarshal(byteValue.([]byte), comic)
	return comic, nil
}

func (d *Dao) GetComicInfoWithCache(ctx context.Context, Channel int, SourceID int) (*model.Comic, error) {
	comic, err := d.CacheGetComicInfo(ctx, Channel, SourceID)
	if nil == comic {
		comic, err = d.GetComicInfo(ctx, Channel, SourceID)
		if err != nil {
			return nil, err
		}
		// 查询不到时，防止缓存穿透，也存储
		err = d.CacheSetComicInfo(ctx, comic)
		if err != nil {
			return nil, err
		}
	} else {
		// 处理缓存穿透
		if 0 == comic.ID {
			return nil, &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND, Message: "漫画不存在!"}
		}
	}
	if nil == comic {
		return nil, &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND, Message: "漫画不存在"}
	}
	return comic, err
}
