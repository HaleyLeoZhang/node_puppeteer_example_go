package po

import "time"

// ----------------------------------------------------------------------
// 基础模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Model struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Status    uint8     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
