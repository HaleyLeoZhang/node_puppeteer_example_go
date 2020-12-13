package comic

import (
	"context"
	"fmt"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
	"github.com/jinzhu/gorm"
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
	constant2 "node_puppeteer_example_go/common/constant"
	"node_puppeteer_example_go/common/model/po"
)

func (d *Dao) GetComicList(ctx context.Context, where map[string]interface{}, attr map[string]interface{}) (res []*po.Comic, err error) {
	res = make([]*po.Comic, 0)
	comicInfo := &po.Comic{}

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

func (d *Dao) GetComicInfo(ctx context.Context, Channel int, SourceID int) (comicInfo *po.Comic, err error) {
	comicInfo = &po.Comic{}

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
	prefix := constant2.CACHE_PREFIX_TYPE_FRONT
	namespace := fmt.Sprintf("page_detail_comic_%v_%v", Channel, SourceID)
	version := "v1"
	suffix := constant2.CACHE_SUFFIX_TYPE_CACHED
	cacheKey := fmt.Sprintf("%v:%v:%v:%v", prefix, namespace, version, suffix)
	return cacheKey
}

func (d *Dao) CacheSetComicInfo(ctx context.Context, comic *po.Comic) error {
	cacheKey := cacheKeyGetComicInfo(comic.Channel, comic.SourceID)
	byteValue, err := easyjson.Marshal(comic)
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

func (d *Dao) CacheGetComicInfo(ctx context.Context, Channel int, SourceID int) (*po.Comic, error) {
	cacheKey := cacheKeyGetComicInfo(Channel, SourceID)
	conn := d.redis.Get()
	defer conn.Close()
	byteValue, err := conn.Do("GET", cacheKey)
	if err != nil {
		return nil, err
	}
	comic := &po.Comic{}
	if nil == byteValue {
		return nil, nil
	}
	easyjson.Unmarshal(byteValue.([]byte), comic)
	return comic, nil
}

func (d *Dao) GetComicInfoWithCache(ctx context.Context, Channel int, SourceID int) (*po.Comic, error) {
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