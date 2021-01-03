package cache

import (
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
	"github.com/gomodule/redigo/redis"
)

type Dao struct {
	Redis *redis.Pool
}

func New(cfg *xredis.Config) *Dao {
	var err error

	d := &Dao{}
	d.Redis, err = xredis.NewPool(cfg)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *Dao) Close() {
	err := d.Redis.Close()
	if err != nil {
		xlog.Errorf("cache.redis.Err(%+v)", err)
	}
}

// --------------------------------
// 漫画基础信息缓存
// --------------------------------
//func cacheKeyGetComicInfo(id int, relatedId int) string {
//	prefix := constant2.CACHE_PREFIX_TYPE_FRONT
//	namespace := fmt.Sprintf("page_detail_comic_%v_%v", id, relatedId)
//	version := "v2"
//	suffix := constant2.CACHE_SUFFIX_TYPE_CACHED
//	cacheKey := fmt.Sprintf("%v:%v:%v:%v", prefix, namespace, version, suffix)
//	return cacheKey
//}
//
//func (d *Dao) CacheSetComicInfo(ctx context.Context, comic *po.Comic) error {
//	cacheKey := cacheKeyGetComicInfo(comic.Id, comic.RelatedId)
//	byteValue, err := easyjson.Marshal(comic)
//	if nil != err {
//		err = errors.WithStack(err)
//		return err
//	}
//	ttl := 300 // 缓存时间，单位，秒
//
//	conn := d.redis.Get()
//	defer conn.Close()
//	_, err = conn.Do("SET", cacheKey, byteValue, "EX", ttl)
//	if err != nil {
//		err = errors.WithStack(err)
//		return err
//	}
//	return err
//}
//
//func (d *Dao) CacheGetComicInfo(ctx context.Context, Channel int, SourceId int) (*po.Comic, error) {
//	cacheKey := cacheKeyGetComicInfo(Channel, SourceId)
//	conn := d.redis.Get()
//	defer conn.Close()
//	byteValue, err := conn.Do("GET", cacheKey)
//	if err != nil {
//		return nil, err
//	}
//	comic := &po.Comic{}
//	if nil == byteValue {
//		return nil, nil
//	}
//	easyjson.Unmarshal(byteValue.([]byte), comic)
//	return comic, nil
//}
//
//func (d *Dao) GetComicInfoWithCache(ctx context.Context, Channel int, SourceId int) (*po.Comic, error) {
//	comic, err := d.CacheGetComicInfo(ctx, Channel, SourceId)
//	if nil == comic {
//		comic, err = d.GetComicInfo(ctx, Channel, SourceId)
//		if err != nil {
//			return nil, err
//		}
//		// 查询不到时，防止缓存穿透，也存储
//		err = d.CacheSetComicInfo(ctx, comic)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		// 处理缓存穿透
//		if 0 == comic.Id {
//			return nil, &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND, Message: "漫画不存在!"}
//		}
//	}
//	if nil == comic {
//		return nil, &xgin.BusinessError{Code: xgin.HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND, Message: "漫画不存在"}
//	}
//	return comic, err
//}
