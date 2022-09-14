package constant

const (
	COMIC_LIST_PAGE_SIZE        = 20
	COMIC_LIST_CACHE_KEY_TPL    = "cm:comic_list:page:%v:v1"
	COMIC_LIST_CACHE_KEY_METRIC = "comic_list"
	COMIC_LIST_CACHE_TTL_MIN    = 5  // 单位，秒
	COMIC_LIST_CACHE_TTL_MAX    = 10 // 单位，秒
)
