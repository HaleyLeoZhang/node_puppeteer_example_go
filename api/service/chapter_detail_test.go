package service

import (
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_ChapterDetail(t *testing.T) {
	param := &model.ChapterDetailParam{
		Id: ctx.Value("chapter_id").(int),
	}
	res, err := svr.ChapterDetail(ctx, param)
	if nil != err {
		t.Fatalf("Err: %v", err)
		return
	}
	raw, _ := json.Marshal(res)
	t.Logf("res %+v param %v", string(raw), param)
}
