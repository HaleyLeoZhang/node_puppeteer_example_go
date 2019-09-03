package models

// ----------------------------------------------------------------------
// 漫画章节列表-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Page struct {
	ID        string `json:"id"`
	Channel   string `json:"channel"`
	ComicID   string `json:"comic_id"`
	Sequence  string `json:"sequence"`
	Name      string `json:"name"`
	Link      string `json:"link"`
	Progress  string `json:"progress"`
	IsDeleted string `json:"is_deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
