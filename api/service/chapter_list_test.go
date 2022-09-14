package service

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	commonservice "github.com/HaleyLeoZhang/node_puppeteer_example_go/common/service"
	"testing"
)

func TestService_ChapterList(t *testing.T) {
	type fields struct {
		commonService *commonservice.Service
	}
	type args struct {
		ctx   context.Context
		param *model.ChapterListParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *model.ChapterListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ctx: ctx,
				param: &model.ChapterListParam{
					ComicId: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svr.ChapterList(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChapterList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			raw, _ := json.Marshal(gotRes)
			t.Logf("%+v", string(raw))
		})
	}
}
