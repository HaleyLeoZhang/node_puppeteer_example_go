package comic

import (
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"node_puppeteer_example_go/api/conf"
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
	d.redis, err = xredis.NewPool(cfg.Redis)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *Dao) Close() {
	err := d.redis.Close()
	if err != nil {
		xlog.Errorf("CloseComicDao.redis.Err(%+v)", err)
	}
	err = d.db.Close()
	if err != nil {
		xlog.Errorf("CloseComicDao.db.Err(%+v)", err)
	}
}
