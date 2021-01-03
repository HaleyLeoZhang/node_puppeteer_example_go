package po

type Supplier struct {
	Model
	RelatedId   int    `json:"related_id"`
	Channel     uint8  `json:"channel"`
	SourceId    string `json:"source_id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Intro       string `json:"intro"`
	MaxSequence int    `json:"max_sequence"`
	Ext1        string `json:"ext1"`
	Ext2        string `json:"ext2"`
	Ext3        string `json:"ext3"`
	Weight      int    `json:"weight"`
}

//数据表---必需
func (s *Supplier) TableName() string {
	return "supplier"
}
