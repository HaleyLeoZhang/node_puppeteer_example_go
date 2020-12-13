package po

// ----------------------------------------------------------------------
// 漫画图片模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type ComicImage struct {
	Model
	PageID   int    `json:"page_id"`
	Sequence int    `json:"sequence"`
	Src      string `json:"src"`
	//Progress string `json:"progress"` // 暂时不需要
}

//数据表---必需
func (c *ComicImage) TableName() string {
	return "comic_images"
}
