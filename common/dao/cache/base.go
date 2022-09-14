package cache

import (
	"context"
	"github.com/HaleyLeoZhang/go-component/driver/xmetric"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/model/bo"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/common/util"
	"github.com/pkg/errors"
)

// 查询单个缓存，获取不到就生成缓存
func (d *Dao) getOrSetCache(ctx context.Context, option *bo.SingleCacheOption) (res []byte, err error) {
	if err = option.Check(); err != nil {
		return
	}

	conn := d.Redis.Get()
	defer conn.Close()
	// Part 1 查询缓存
	resInterface, err := conn.Do("GET", option.CacheKey)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if resInterface != nil {
		// - Metric统计
		if option.MetricName != "" {
			xmetric.MetricHit.WithLabelValues(option.MetricName).Inc() // 使用现成指标
		}
		res = resInterface.([]byte)
		return
	}

	res, err = option.CallableToGetValue()
	if err != nil {
		return
	}
	var ttl = option.TtlMin
	if option.TtlMin > 0 && option.TtlMax > option.TtlMin {
		ttl = util.RandInt(option.TtlMin, option.TtlMax) // 缓存时间，单位，秒
	}
	// 生成缓存
	_, err = conn.Do("SET", option.CacheKey, res, "EX", ttl)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// - Metric统计
	if option.MetricName != "" {
		xmetric.MetricMiss.WithLabelValues(option.MetricName).Inc() // 使用现成指标
	}
	return
}
