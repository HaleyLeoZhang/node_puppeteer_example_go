package comic

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/component/driver/db"
	ownredis "node_puppeteer_example_go/component/driver/redis"
)

type Dao struct {
	db    *gorm.DB
	redis *redis.Pool
}

func New(cfg *conf.Config) *Dao {
	var err error

	d := &Dao{}
	if d.db, err = db.New(cfg.DB); err != nil {
		panic(err)
	}
	d.redis, err = ownredis.NewPool(cfg.Redis)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *Dao) Close() {
	d.redis.Close()
	d.db.Close()
}

func (d *Dao) Ping(ctx context.Context) error {
	return nil
}
