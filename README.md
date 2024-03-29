## 序言

漫画爬虫配套的 `API` 服务

* `go mod` 管理包
* `Makefile` 集成命令
* `api/build/main.go` 入口文件

### 性能表现

本次压测因条件限制，在业务宿主机(2c/2G)上直接进行的压测

| 本期 `API` 服务 - 配置指标 | 值 |
| --- | --- |
|  Unix 环境 - CPU最大资源利用率 | 30% |
| 内存最大允许占用 | 200MB |

最大支持 `QPS` 在 `3170` 左右  
最大支持 `150` 个并发连接数左右  

其实本次压测的 `API` 基于本地缓存 `singlefight` 做的本地 `一级接口级缓存` 的

![](doc/readme_performance.jpg)  
`图 1-1 - API性能表现`

作者本次就直接展示下压测前后的业务机环境的资源运用情况了  

压测前 如 `图 02`  
![](doc/before.jpg)  
`图 1-2`

压测时 如 `图 03`  
![](doc/doing.jpg)  
`图 1-3`

可以看到本次瓶颈在 `CPU` 资源上  

#### 二级缓存

`2022-12月` 作者对于漫画首页接口，基于 `Redis` 缓存中间件 做了 `二级缓存`  
使用与前文相同的压测条件

![](doc/wrk_with_cache.jpeg)  
`图 2-1`

各指标表现结果

![](doc/qps_with_cache.png)  
`图 2-2`

![](doc/cache_hit_rate_in_1_minute.png)  
`图 2-3`

![](doc/request_counter_in_1_minute.png)  
`图 2-4`

![](doc/response_ms_in_1_minute.png)  
`图 2-5`

> 相关文章

有兴趣的读者朋友可以通过 [http://www.hlzblog.top/article/74.html](http://www.hlzblog.top/article/74.html)  
分析下平时压测的 `API` 瓶颈是哪里

> 运行前要求

生成好 `app.yaml`  

~~~bash
make ini
~~~

### 使用步骤

> 调试阶段

~~~bash
make
~~~

如果修改了 `Model` 层数据结构，请重新生成 `easyjson`

~~~bash
make json
~~~

> 生成并运行应用

~~~bash
make clean && make build && make run
~~~

> 格式化代码

~~~bash
make tool
~~~

> 优雅关闭

会监听指定信号 ctrl+c 、kill 进程Pid，关闭各种链接后，慢慢退出

~~~bash
- 请不要用 kill -9 程序监听不到退出
- 调试的时候 kill 目标请杀掉对应 tmp 进程即可 
~~~

###### 注意事项

> Goland

在 `goland` 的 `setting` 里设置启用`Go Modules`

~~~bash
goland Preference->Go->Go Modules(vgo) -> Enable Go Modules(vgo)intergration
~~~