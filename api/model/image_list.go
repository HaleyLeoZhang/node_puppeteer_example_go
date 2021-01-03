package model

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ImageListParam struct {
	ChapterId int `form:"chapter_id" binding:"required,gte=1"`
}
type ImageListResponse struct {
	List []*ImageListResponseItem `json:"list"`
}
type ImageListResponseItem struct {
	Sequence  int    `json:"sequence"`
	SrcOrigin string `json:"src_origin"`
	SrcOwn    string `json:"src_own"`
}
