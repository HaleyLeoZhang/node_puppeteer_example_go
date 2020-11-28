package vo

import "node_puppeteer_example_go/common/model/po"

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type PageListParam struct {
	Channel  int `form:"channel" binding:"gte=1"`
	SourceId int `form:"source_id" binding:"gte=0"`
}
type PageListResponse struct {
	List []*po.ComicPage `json:"list"`
}
