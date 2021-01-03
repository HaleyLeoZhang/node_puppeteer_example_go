package po

type SupplierImage struct {
	Model
	RelatedId int    `json:"related_id"`
	Sequence  int    `json:"sequence"`
	SrcOrigin string `json:"src_origin"`
	SrcOwn    string `json:"src_own"`
	Progress  int    `json:"progress"`
}

//数据表---必需
func (s *SupplierImage) TableName() string {
	return "supplier_image"
}
