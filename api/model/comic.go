package model

// ----------------------------------------------------------------------
// 漫画基础信息模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Comic struct {
	*Model
	Channel     string `json:"channel"`
	SourceID    string `json:"source_id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Intro       string `json:"intro"`
	MaxSequence string `json:"max_sequence"`
	Weight      string `json:"weight"`
	Tag         string `json:"tag"`
}

//数据表---必需
func (Comic) TableName() string {
	return "comics"
}

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ComicListParam struct {
	Page int `form:"page" binding:"gte=1"`
}

type ComicListResponse struct {
	List *[]Comic `json:"list"`
}
