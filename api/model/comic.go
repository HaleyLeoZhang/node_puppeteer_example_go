package model

// ----------------------------------------------------------------------
// 漫画基础信息模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Comic struct {
	*Model
	Channel     int    `json:"channel"`
	SourceID    int    `json:"source_id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Intro       string `json:"intro"`
	MaxSequence int    `json:"max_sequence"`
	Weight      int    `json:"weight"`
	Tag         int    `json:"tag"`
}

//数据表---必需
func (c *Comic) TableName() string {
	return "comics"
}

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ComicListParam struct {
	Page int `form:"page" binding:"gte=1"`
}

type ComicListResponse struct {
	List []*Comic `json:"list"`
}
