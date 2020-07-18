package model

// ----------------------------------------------------------------------
// 基础模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	//IsDeleted string `json:"is_deleted"` // 暂时用不上
	//UpdatedAt string `json:"updated_at"` // 暂时用不上
	//CreatedAt string `json:"created_at"` // 暂时用不上
}
