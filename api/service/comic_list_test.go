package service

import (
	"context"
	"encoding/json"
	"github.com/HaleyLeoZhang/node_puppeteer_example_go/api/model"
	"testing"
)

func TestService_ComicList(t *testing.T) {
	type args struct {
		ctx   context.Context
		param *model.ComicListParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ctx: ctx,
				param: &model.ComicListParam{
					Page: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svr.ComicList(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComicList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			raw, _ := json.Marshal(gotRes)
			t.Logf("%+v", string(raw))
		})
	}
}
