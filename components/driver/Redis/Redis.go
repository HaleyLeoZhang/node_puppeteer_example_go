package Redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Conf struct {
	*pool.Config `yaml:"pool"`

	Name         string         `yaml:"name"` // redis name, for trace
	Proto        string         `yaml:"proto"`
	Addr         string         `yaml:"addr"`
	Auth         string         `yaml:"auth"`
	DialTimeout  xtime.Duration `yaml:"dialTimeout"`
	ReadTimeout  xtime.Duration `yaml:"readTimeout"`
	WriteTimeout xtime.Duration `yaml:"writeTimeout"`
	DB           int            `yaml:"db"`
	SlowLog      xtime.Duration `yaml:"slowLog"`
}

type Pool struct {
	 //
}
// Setup Initialize the Redis instance
func New(c Conf) error {
	RedisConn := &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}
