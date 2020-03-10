## 序言

漫画爬虫配套的 `API` 服务  

* 基于 [go-gin-example](https://github.com/eddycjy/go-gin-example/blob/master/README_ZH.md) 改造  
* `go mod`管理包
* `Makefile` 集成命令
* [查看接口文档](http://api_puppeteer.doc.hlzblog.top/)  

> 运行前要求

生成好 `app.ini`  

~~~bash
cp conf/app.ini.example conf/app.ini
~~~

### 使用步骤

> 生成并运行应用

~~~bash
make build
./puppeteer.hlzblog.top
~~~

> 格式化代码

~~~bash
make tool
~~~