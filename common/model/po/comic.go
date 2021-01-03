package po

// ----------------------------------------------------------------------
// 漫画基础信息模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Comic struct {
	Model
	RelatedId int    `json:"related_id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	Intro     string `json:"intro"`
	Weight    int    `json:"weight"`
	Tag       int    `json:"tag"`
	Method    int    `json:"method"`
}

//数据表---必需
func (c *Comic) TableName() string {
	return "comic"
}
