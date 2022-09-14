package service

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_ChapterDetail(t *testing.T) {

	type args struct {
		ctx   context.Context
		param *model.ChapterDetailParam
	}
	tests := []struct {
		name    string
		args    args
		wantRes *model.ChapterDetailResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ctx: ctx,
				param: &model.ChapterDetailParam{
					Id: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svr.ChapterDetail(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChapterDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			raw, _ := json.Marshal(gotRes)
			t.Logf("%+v", string(raw))
		})
	}
}
