package po

// ----------------------------------------------------------------------
// 漫画基础信息模型
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Comic struct {
	Model
	RelatedId int    `gorm:"column:related_id"`
	Name      string `gorm:"column:name"`
	Pic       string `gorm:"column:pic"`
	Intro     string `gorm:"column:intro"`
	Weight    int    `gorm:"column:weight"`
	Tag       int    `gorm:"column:tag"`
	Method    int    `gorm:"column:method"`
}

//数据表---必需
func (c *Comic) TableName() string {
	return "comic"
}
