package po

type SupplierChapter struct {
	Model
	RelatedId int    `json:"related_id"`
	Sequence  int    `json:"sequence"`
	Name      string `json:"name"`
}

//数据表---必需
func (s *SupplierChapter) TableName() string {
	return "supplier_chapter"
}
