package vo

import "node_puppeteer_example_go/common/model/po"

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ImageListParam struct {
	PageId int `form:"page_id" binding:"required,gte=1"`
}
type ImageListResponse struct {
	List []*po.ComicImage `json:"list"`
}
