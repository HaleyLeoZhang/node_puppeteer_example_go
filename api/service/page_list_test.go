package service

import (
	"encoding/json"
	"node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_PageList(t *testing.T) {
	param := &model.PageListParam{
		Channel:  ctx.Value("channel").(int),
		SourceId: ctx.Value("source_id").(int),
	}
	res, err := svr.PageList(ctx, param)
	if nil != err {
		t.Fatalf("Err: %v", err)
	} else {
		raw, _ := json.Marshal(res)
		t.Logf("res %+v param %v", string(raw), param)
	}
}
