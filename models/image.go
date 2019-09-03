package models

// ----------------------------------------------------------------------
// 漫画章节对应图片列表-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Image struct {
	ID        string `json:"id"`
	PageID    string `json:"page_id"`
	Sequence  string `json:"sequence"`
	Src       string `json:"src"`
	Progress  string `json:"progress"`
	IsDeleted string `json:"is_deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
