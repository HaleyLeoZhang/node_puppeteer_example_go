package comic

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/conf"
	"node_puppeteer_example_go/component/driver/db"
	"node_puppeteer_example_go/component/driver/ownlog"
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
	err := d.redis.Close()
	if err != nil {
		ownlog.Errorf("CloseComicDao.redis.Err(%+v)", err)
	}
	err = d.db.Close()
	if err != nil {
		ownlog.Errorf("CloseComicDao.db.Err(%+v)", err)
	}
}
