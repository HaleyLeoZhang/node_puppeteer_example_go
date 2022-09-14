package bo

import (
	"fmt"
	"github.com/pkg/errors"
)

// 请求查询或生成单条缓存数据

type SingleCacheOption struct {
	// 【必填】缓存的名称
	CacheKey string
	// 缓存打点指标名。如果为空，则不打点
	MetricName string
	// 【必填】通过回调函数，在没有命中缓存的时候，生成查询结果
	CallableToGetValue func() ([]byte, error)
	// 缓存的最小，最大时间，单位秒
	//		如果最小最大时间区间都设置了，会以范围区间，随机生成ttl
	//  	如果只设置了 TtlMin ，则是固定时长
	TtlMin int // 【必填】
	TtlMax int
}

func (option *SingleCacheOption) Check() (err error) {
	if option.TtlMin <= 0 {
		err = errors.WithStack(fmt.Errorf("缓存 TtlMin 必须大于0"))
		return
	}
	if option.TtlMax > 0 && option.TtlMax < option.TtlMin {
		err = errors.WithStack(fmt.Errorf("缓存 TtlMax 必须大于 TtlMin"))
		return
	}
	if len(option.CacheKey) == 0 {
		err = errors.WithStack(fmt.Errorf("缓存 CacheKey 不能为空"))
		return
	}
	if option.CallableToGetValue == nil {
		err = errors.WithStack(fmt.Errorf("缓存 CallableToGetValue 必须设置"))
		return
	}
	return
}
