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
		xlog.Errorf("cache redis Err(%+v)", err)
	}
}
