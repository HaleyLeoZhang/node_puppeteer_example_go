package po

type SupplierImage struct {
	Model
	RelatedId int    `gorm:"column:related_id"`
	Sequence  int    `gorm:"column:sequence"`
	SrcOrigin string `gorm:"column:src_origin"`
	SrcOwn    string `gorm:"column:src_own"`
	Progress  int    `gorm:"column:progress"`
}

//数据表---必需
func (s *SupplierImage) TableName() string {
	return "supplier_image"
}
