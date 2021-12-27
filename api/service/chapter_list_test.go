package service

import (
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_PageList(t *testing.T) {
	param := &model.ChapterListParam{
		ComicId: ctx.Value("comic_id").(int),
	}
	res, err := svr.ChapterList(ctx, param)
	if nil != err {
		t.Fatalf("Err: %v", err)
		return
	}
	raw, _ := json.Marshal(res)
	t.Logf("res %+v param %v", string(raw), param)
}
