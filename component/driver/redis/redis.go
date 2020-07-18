package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//  DB使用默认0库
type Config struct {
	Name         string        `yaml:"name"` // 用于 Trace 识别
	Proto        string        `yaml:"proto"`
	Addr         string        `yaml:"addr"`
	Auth         string        `yaml:"auth"`
	DialTimeout  time.Duration `yaml:"dialTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	SlowLog      time.Duration `yaml:"slowLog"`
	Pool         PoolConfig    `yaml:"pool"`
}

type PoolConfig struct {
	MaxActive   int           `yaml:"maxActive"`
	MaxIdle     int           `yaml:"maxIdle"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

func NewPool(c *Config) (*redis.Pool, error) {
	RedisConn := &redis.Pool{
		MaxActive:   c.Pool.MaxActive,
		MaxIdle:     c.Pool.MaxIdle,
		IdleTimeout: c.Pool.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Proto, c.Addr)
			if err != nil {
				return nil, err
			}
			if c.Auth != "" {
				if _, err := conn.Do("AUTH", c.Auth); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}

	return RedisConn, nil
}
