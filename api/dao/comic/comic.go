package comic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/constant"
	"node_puppeteer_example_go/api/model"
	"node_puppeteer_example_go/component/driver/db"
)

func (d *Dao) GetComicList(ctx context.Context, page int, size int, maps map[string]interface{}) (*[]model.Comic, error) {
	comicList := make([]model.Comic, 0)
	offset, size := db.GetPageInfo(page, size)

	maps["is_deleted"] = constant.TABLE_BASE_IS_DELETED_NO

	err := d.db.Where(maps).Offset(offset).Order("weight DESC").Limit(size).Find(&comicList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return &comicList, err
	}

	return &comicList, nil
}

func (d *Dao) GetComicInfo(ctx context.Context, Channel int, SourceID int) (*model.Comic, error) {
	comicInfo := &model.Comic{}

	maps := make(map[string]interface{})
	maps["channel"] = Channel
	maps["source_id"] = SourceID

	err := d.db.Where(maps).First(&comicInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("error %+v", err)
		return nil, err
	}

	return comicInfo, nil
}

// --------------------------------
// 漫画基础信息缓存
// --------------------------------
func cacheKeyGetComicInfo(Channel int, SourceID int) string {
	prefix := constant.CACHE_PREFIX_TYPE_FRONT
	namespace := fmt.Sprintf("page_detail_comic_%v_%v", Channel, SourceID) // https://sh-api.shihuo.cn/app_swoole_zone/sub
	version := "v1"
	suffix := constant.CACHE_SUFFIX_TYPE_CACHED
	cacheKey := fmt.Sprintf("%v:%v:%v:%v", prefix, namespace, version, suffix)
	return cacheKey
}

func (d *Dao) CacheSetComicInfo(ctx context.Context, comic *model.Comic) error {
	cacheKey := cacheKeyGetComicInfo(comic.Channel, comic.SourceID)
	byteValue, err := json.Marshal(comic)
	if nil != err {
		fmt.Printf("error %+v", err)
		return err
	}
	ttl := 300 // 缓存时间，单位，秒

	conn := d.redis.Get()
	defer conn.Close()
	_, err = conn.Do("SET", cacheKey, byteValue, "EX", ttl)
	if err != nil {
		fmt.Printf("error %+v", err)
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
		return nil, err
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
			return nil, errors.New("漫画不存在!") // 用 感叹号区分错误类型
		}
	}
	if nil == comic {
		return nil, errors.New("漫画不存在")
	}
	return comic, err
}
