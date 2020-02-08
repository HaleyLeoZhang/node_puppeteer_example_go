package caches

import (
	"encoding/json"
	"strconv"
	"strings"

	"node_puppeteer_example_go/models"
	"node_puppeteer_example_go/pkg/e"
	"node_puppeteer_example_go/pkg/gredis"
	"node_puppeteer_example_go/pkg/logging"
)

type ComicPage struct {
	Channel  int
	SourceID int
}

func (c *ComicPage) getKeyName() string {
	keys := []string{}

	keys = append(keys, e.CACHE_COMIC_PAGE_LIST)
	keys = append(keys, strconv.Itoa(c.Channel))
	keys = append(keys, strconv.Itoa(c.SourceID))

	return strings.Join(keys, e.DELEMITER_CACHE)
}

func (c *ComicPage) Get() ([]*models.ComicPages, error) {
	key := c.getKeyName()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			// logging.Info("在走缓存了")
			var PageList []*models.ComicPages
			json.Unmarshal(data, &PageList)
			return PageList, nil
		}
	}
	return nil, nil
}

func (c *ComicPage) Save(PageList []*models.ComicPages) {
	key := c.getKeyName()
	ttl := 60
	gredis.Set(key, PageList, ttl)
}
