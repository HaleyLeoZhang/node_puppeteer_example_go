package model

// ----------------------------------------------------------------------
// 漫画图片模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type ComicImage struct {
	*Model
	PageID   int    `json:"page_id"`
	Sequence int    `json:"sequence"`
	Src      string `json:"src"`
	//Progress string `json:"progress"` // 暂时不需要
}

//数据表---必需
func (ComicImage) TableName() string {
	return "comic_images"
}

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ImageListParam struct {
	PageId int `form:"page_id" binding:"required,gte=1"`
}
type ImageListResponse struct {
	List *[]ComicImage `json:"list"`
}
