package model

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ChapterListParam struct {
	ComicId int `form:"comic_id" binding:"gte=1"`
}
type ChapterListResponse struct {
	List []*ChapterListResponseItem `json:"list"`
}
type ChapterListResponseItem struct {
	Id       int    `json:"id"`
	Sequence int    `json:"sequence"`
	Name     string `json:"name"`
}
