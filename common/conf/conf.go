package conf

import (
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
)

// Config struct
type Config struct {
	DB    *db.Config     `yaml:"db" json:"db"`
	Redis *xredis.Config `yaml:"redis" json:"redis"`
}
