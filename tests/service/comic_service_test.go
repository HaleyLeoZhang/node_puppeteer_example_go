package test_service

import (
	"node_puppeteer_example_go/comic_service"
    "testing"
)

// go test -v

func newService() *email_service.Email{} {
	return &email_service.Email{}
}

func TestGetArea(t *testing.T) {
    service := newService()

    data := new(map[string]interface{})
    data["title"] = "测试"
    data["content"] = "文本"
    data["sender_name"] = "云天河测试"
    data["receiver"] = "229270575@qq.com"
    data["receiver_name"] = "沐临风"

    error := service.DoCreate(data)
    if error != nil {
        t.Error("测试失败")
    }else{
        t.Error("success")
    }
}
