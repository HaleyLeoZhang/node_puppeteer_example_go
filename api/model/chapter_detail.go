package model

// 验证器规则 https://blog.csdn.net/guyan0319/article/details/105918559/
type ChapterDetailParam struct {
	Id int `form:"id" binding:"gte=1"`
}
type ChapterDetailResponse struct {
	Chapter     *ChapterDetailResponseChapter `json:"chapter"`
	NextChapter *ChapterDetailResponseChapter `json:"next_chapter"`
	Comic       *ChapterDetailResponseComic   `json:"comic"`
}

type ChapterDetailResponseChapter struct {
	Id       int    `json:"id"`
	Sequence int    `json:"sequence"`
	Name     string `json:"name"`
}

type ChapterDetailResponseComic struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
}
