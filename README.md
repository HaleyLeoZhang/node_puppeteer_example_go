# 序言
这里将会用 [golang](https://golang.org/) 中的 [gin](https://www.yoytang.com/go-gin-doc.html) 框架 吐出图片地址数据  

[查看接口文档](http://api_puppeteer.doc.hlzblog.top/)  

## 基于框架
该框架使用请看这里  
[go-gin-example](https://github.com/HaleyLeoZhang/node_puppeteer_example_go/blob/master/README_ZH.md)   

~~~bash
cp conf/app.ini.example conf/app.ini  
~~~

## 工具

[json转go结构体](https://www.sojson.com/json/json2go.html)  
请注意json中的数据类型  

## 常用功能

> 生成应用

~~~bash
go build -o puppeteer.hlzblog.top main.go
~~~

或者  

~~~bash
go build -v .
~~~

> 格式化代码

~~~bash
gofmt -w .
~~~