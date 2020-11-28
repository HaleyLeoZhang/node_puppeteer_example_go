package vo

import "node_puppeteer_example_go/common/model/po"

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ComicListParam struct {
	Page int `form:"page" binding:"gte=1"`
}

type ComicListResponse struct {
	List []*po.Comic `json:"list"`
}
