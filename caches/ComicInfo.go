package caches

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HaleyLeoZhang/node_puppeteer_example_go/models"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/e"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/gredis"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/pkg/logging"
)

type ComicInfo struct {
	Channel int
	ComicID int
}

func (c *ComicInfo) getKeyName() string {
	keys := []string{}

	keys = append(keys, e.CACHE_COMIC_INFO)
	keys = append(keys, strconv.Itoa(c.Channel))
	keys = append(keys, strconv.Itoa(c.ComicID))

	return strings.Join(keys, e.DELEMITER_CACHE)
}

func (c *ComicInfo) Get() (*models.Comics, error) {
	key := c.getKeyName()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			// logging.Info("在走缓存了")
			var OneComic *models.Comics
			json.Unmarshal(data, &OneComic)
			return OneComic, nil
		}
	}
	return nil, nil
}

func (c *ComicInfo) Save(OneComic *models.Comics) {
	key := c.getKeyName()
	ttl := 300
	gredis.Set(key, OneComic, ttl)
}
