package service

import (
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_ComicList(t *testing.T) {
	param := &model.ComicListParam{
		Page: ctx.Value("page").(int),
	}
	res, err := svr.ComicList(ctx, param)
	if nil != err {
		t.Fatalf("Err: %+v", err)
		return
	}
	raw, _ := json.Marshal(res)
	t.Logf("res %+v param %v", string(raw), param)
}
