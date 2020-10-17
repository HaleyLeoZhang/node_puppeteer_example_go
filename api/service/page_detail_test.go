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

func Benchmark_PageDetail(t *testing.B) {
	t.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	t.StartTimer() //重新开始时间
	param := &model.PageDetailParam{
		PageId: ctx.Value("page_id").(int),
	}

	t.N=1234 //自定义执行1234次

	for i := 0; i < t.N; i++ {
		res, err := svr.PageDetail(ctx, param)
		if nil != err {
			t.Fatalf("Err: %v", err)
		} else {
			raw, _ := json.Marshal(res)
			t.Logf("res %+v param %v", string(raw), param)
		}
	}

}