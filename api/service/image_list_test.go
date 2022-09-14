package service

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	commonservice "github.com/HaleyLeoZhang/node_puppeteer_example_go/common/service"
	"testing"
)

func TestService_ImageList(t *testing.T) {
	type fields struct {
		commonService *commonservice.Service
	}
	type args struct {
		ctx   context.Context
		param *model.ImageListParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *model.ImageListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ctx: ctx,
				param: &model.ImageListParam{
					ChapterId: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svr.ImageList(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			raw, _ := json.Marshal(gotRes)
			t.Logf("%+v", string(raw))
		})
	}
}
