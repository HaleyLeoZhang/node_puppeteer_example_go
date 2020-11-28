package service

import (
	"encoding/json"
	"node_puppeteer_example_go/common/model/vo"
	"testing"
)

func TestService_ComicList(t *testing.T) {
	param := &vo.ComicListParam{
		Page: ctx.Value("page").(int),
	}
	res, err := svr.ComicList(ctx, param)
	if nil != err {
		t.Fatalf("Err: %v", err)
	} else {
		raw, _ := json.Marshal(res)
		t.Logf("res %+v param %v", string(raw), param)
	}
}
