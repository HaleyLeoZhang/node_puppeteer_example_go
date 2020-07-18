package model

// ----------------------------------------------------------------------
// 漫画章节模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type ComicPage struct {
	*Model
	Channel  int    `json:"channel"`
	SourceId int    `json:"source_id"`
	Sequence int    `json:"sequence"`
	Name     string `json:"name"`
	//Link     string `json:"link"`
	//Progress int    `json:"progress"`
}

//数据表---必需
func (ComicPage) TableName() string {
	return "comic_pages"
}

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type PageListParam struct {
	Channel  int `form:"channel" binding:"required,gte=0"`
	SourceId int `form:"source_id" binding:"required,gte=0"`
}
type PageListResponse struct {
	List *[]ComicPage `json:"list"`
}

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type PageDetailParam struct {
	PageId int `form:"page_id" binding:"required,gte=1"`
}
type PageDetailResponse struct {
	Page     *ComicPage `json:"page"`
	NextPage *ComicPage `json:"next_page"`
	Comic    *Comic     `json:"comic"`
}
