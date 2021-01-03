package service

import (
	"encoding/json"
	"node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_ImageList(t *testing.T) {

	param := &model.ImageListParam{
		ChapterId: ctx.Value("chapter_id").(int),
	}
	res, err := svr.ImageList(ctx, param)
	if nil != err {
		t.Fatalf("Err: %+v", err)
		return
	}
	raw, _ := json.Marshal(res)
	t.Logf("res %+v param %v", string(raw), param)
}
