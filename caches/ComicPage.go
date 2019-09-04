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

type ComicPage struct {
	Channel int
	ComicID int
}

func (c *ComicPage) getKeyName() string {
	keys := []string{}

	keys = append(keys, e.CACHE_COMICPAGE)
	keys = append(keys, strconv.Itoa(c.Channel))
	keys = append(keys, strconv.Itoa(c.ComicID))

	return strings.Join(keys, ":")
}

func (c *ComicPage) Get() ([]*models.Pages, error) {
	key := c.getKeyName()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			// logging.Info("在走缓存了")
			var PageList []*models.Pages
			json.Unmarshal(data, &PageList)
			return PageList, nil
		}
	}
	return nil, nil
}

func (c *ComicPage) Save(PageList []*models.Pages) {
	key := c.getKeyName()
	ttl := 60
	gredis.Set(key, PageList, ttl)
}
