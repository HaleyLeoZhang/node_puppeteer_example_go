package comic

import (
"context"

)

//Dao struct user of color egg Dao.
type Dao struct {
	db         *orm.Engine
	redis      *redis.Pool
	httpClient *bm.Client
}

//New .
func New(cfg *conf.Config) (d *Dao) {
	var (
		err error
	)
	d = &Dao{
		httpClient: bm.NewClient(cfg.HttpClient),
	}
	if d.db, err = orm.New(cfg.Drds); err != nil {
		panic(err)
	}
	d.redis = redis.NewPool(cfg.Redis)
	return
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
	d.db.Close()
}

