package DB

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)


type Config struct {
	Type         string        `yaml:"type" json:"type"`
	Server       string        `yaml:"server" json:"server"`
	Port         int         `yaml:"port" json:"port"`
	Database     string        `yaml:"database" json:"database"`
	User         string        `yaml:"user" json:"user"`
	Password     string        `yaml:"password" json:"password"`
	MaxIdleConns int           `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int           `yaml:"maxOpenConns" json:"maxOpenConns"`
	MaxLeftTime  time.Duration `yaml:"maxLeftTime" json:"maxLeftTime"`
}

func New(conf *Config) {
	var err error
	// db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&loc=Local",
		conf.User, conf.Password, conf.Server, conf.Port, conf.Database)
	db, err := gorm.Open(conf.Type, dsn)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return db
}

func Close(db) {
	defer db.Close()
}