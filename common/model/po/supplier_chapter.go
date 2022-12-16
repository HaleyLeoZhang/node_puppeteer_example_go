package po

type SupplierChapter struct {
	Model
	RelatedId int    `gorm:"column:related_id"`
	Sequence  int    `gorm:"column:sequence"`
	Name      string `gorm:"column:name"`
}

//数据表---必需
func (s *SupplierChapter) TableName() string {
	return "supplier_chapter"
}
