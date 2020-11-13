package vo

import "node_puppeteer_example_go/api/model/po"

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type PageDetailParam struct {
	PageId int `form:"page_id" binding:"gte=1"`
}
type PageDetailResponse struct {
	Page     *po.ComicPage `json:"page"`
	NextPage *po.ComicPage `json:"next_page"`
	Comic    *po.Comic     `json:"comic"`
}
