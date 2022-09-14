package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/bo"
	"github.com/pkg/errors"
)

func (d *Dao) ComicListRequestWithCache(ctx context.Context, page int, callable func() ([]byte, error)) (res *model.ComicListResponse, err error) {
	res = &model.ComicListResponse{}
	// -
	resBytes, err := d.getOrSetCache(ctx, &bo.SingleCacheOption{
		CacheKey:           fmt.Sprintf(constant.COMIC_LIST_CACHE_KEY_TPL, page),
		MetricName:         constant.COMIC_LIST_CACHE_KEY_METRIC,
		CallableToGetValue: callable,
		TtlMin:             constant.COMIC_LIST_CACHE_TTL_MIN,
		TtlMax:             constant.COMIC_LIST_CACHE_TTL_MAX,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
