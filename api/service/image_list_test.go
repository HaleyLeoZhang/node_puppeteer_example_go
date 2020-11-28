package service

import (
	"encoding/json"
	"node_puppeteer_example_go/common/model/vo"
	"testing"
)

func TestService_ImageList(t *testing.T) {

	param := &vo.ImageListParam{
		PageId: ctx.Value("page_id").(int),
	}
	res, err := svr.ImageList(ctx, param)
	if nil != err {
		t.Fatalf("Err: %v", err)
	} else {
		raw, _ := json.Marshal(res)
		t.Logf("res %+v param %v", string(raw), param)
	}
}
