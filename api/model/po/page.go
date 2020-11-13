package po

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
	Progress int `json:"progress"`
}

//数据表---必需
func (c *ComicPage) TableName() string {
	return "comic_pages"
}
