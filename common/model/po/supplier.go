package po

type Supplier struct {
	Model
	RelatedId   int    `gorm:"column:related_id"`
	Channel     int    `gorm:"column:channel"`
	SourceId    string `gorm:"column:source_id"`
	Name        string `gorm:"column:name"`
	Pic         string `gorm:"column:pic"`
	Intro       string `gorm:"column:intro"`
	MaxSequence int    `gorm:"column:max_sequence"`
	Ext1        string `gorm:"column:ext1"`
	Ext2        string `gorm:"column:ext2"`
	Ext3        string `gorm:"column:ext3"`
	Weight      int    `gorm:"column:weight"`
}

//数据表---必需
func (s *Supplier) TableName() string {
	return "supplier"
}
