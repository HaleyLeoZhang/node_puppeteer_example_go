package constant

// ----------------------------------------------------------------------
// 缓存公共配置
// ----------------------------------------------------------------------

const (
	CACHE_DELEMITER = ":"
	/**
	 * 命名规范
	 * 前缀:命名空间从大到小:版本号:后缀
	 * - front:comic_list_page_1:v1:cached
	 * - front:comic_info_chnnel_6_source_12006:v2:storage
	 */
	CACHE_PREFIX_TYPE_FRONT   = "front"   // 前缀: 前台接口的缓存
	CACHE_PREFIX_TYPE_BACKEND = "backend" // 前缀: 是后台接口的缓存
	CACHE_SUFFIX_TYPE_CACHED  = "cached"  // 后缀：缓存数据
	CACHE_SUFFIX_TYPE_STORAGE = "storage" // 后缀: 持久化数据
)
