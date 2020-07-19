package service

import (
	"encoding/json"
	"node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_PageDetail(t *testing.T) {
	param := &model.PageDetailParam{
		PageId: ctx.Value("page_id").(int),
	}
	res, err := svr.PageDetail(ctx, param)
	if nil != err {
		t.Fatalf("Err: %v", err)
	} else {
		raw, _ := json.Marshal(res)
		t.Logf("res %+v param %v", string(raw), param)
	}
}
