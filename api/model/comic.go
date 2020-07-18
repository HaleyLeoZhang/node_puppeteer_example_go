package model

// ----------------------------------------------------------------------
// 漫画列表-模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"github.com/jinzhu/gorm"
)

type Comics struct {
	ID       string `json:"id"`
	Channel  string `json:"channel"`
	SourceID string `json:"source_id"`
	Name     string `json:"name"`
	Pic      string `json:"pic"`
	Intro    string `json:"intro"`
	// IsDeleted string `json:"is_deleted"`
	MaxSequence string `json:"max_sequence"`
	Weight      string `json:"weight"`
	Tag         string `json:"tag"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
