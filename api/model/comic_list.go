package model

import "github.com/HaleyLeoZhang/node_puppeteer_example_go/common/constant"

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ComicListParam struct {
	Page int `form:"page" binding:"gte=1"`
}

func (input *ComicListParam) GetPage() int {
	return input.Page
}

func (input *ComicListParam) GetPageSize() int {
	return constant.COMIC_LIST_PAGE_SIZE
}

type ComicListResponse struct {
	List []*ComicListResponseItem `json:"list"`
}
type ComicListResponseItem struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Intro  string `json:"intro"`
	Weight int    `json:"weight"`
	Tag    int    `json:"tag"`

	Supplier *ComicListResponseSupplier `json:"supplier"`
}
type ComicListResponseSupplier struct {
	Id          int `json:"id"`
	MaxSequence int `json:"max_sequence"`
}
