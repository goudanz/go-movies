# Go Movies

> 一个基于golang的爬虫电影站，效果站： [https://go-movies.hezhizheng.com/](https://go-movies.hezhizheng.com/)

![img](https://i.loli.net/2019/12/05/Qzqv4HWoMp2DByi.png)

## 使用安装 
```
# 下载
git clone https://github.com/hezhizheng/go-movies

# 进入目录
cd go-movies

# 生成配置文件(默认使用redis db10的库，可自行修改app.go中的配置)
cp ./config/app.go.backup ./config/app.go

# 启动
go run main.go 
or
# 安装 bee 工具
bee run

# 如安装依赖包失败，请使用代理
export GOPROXY=https://goproxy.io,direct
or
export GOPROXY=https://goproxy.cn,direct

访问
http://127.0.0.1:8899
```

### 开启爬虫
- 直接访问链接http://127.0.0.1:8899/movies-spider(开启定时任务，定时爬取就好)
- 消耗：Windows 下 cup 10% 左右，内存 30mb 左右(爬虫完毕都会降下来) 
- 网络正常的情况下，爬虫完毕耗时大概21分钟左右（存在部分资源爬取失败的情况）

## Tools
- [https://github.com/gocolly/colly](https://github.com/gocolly/colly) 爬虫框架
- 模板引擎：https://github.com/shiyanhui/hero
- 数据库 redis 缓存/持久 [https://github.com/Go-redis/redis](https://github.com/Go-redis/redis)
- 路由 [https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- json解析 jsoniter [github.com/json-iterator/go](github.com/json-iterator/go)
- 跨平台打包：https://github.com/mitchellh/gox
- 静态资源处理：https://github.com/rakyll/statik

## 注意
```
# 修改静态文件/static  、 views/hero 需要先安装包的依赖，执行以下编译命令，更多用法可参考官方redame.md

# https://github.com/rakyll/statik
statik -src=xxxPath/go_movies/static -f 

# https://github.com/shiyanhui/hero
hero -source="./views/hero"
```

## 编译可执行文件(跨平台)
```
# 用法参考 https://github.com/mitchellh/gox
# 生成文件可直接执行
gox -osarch="linux/amd64" # Linux
......

```

## 目录结构参考beego设置

## TODO
- [x] 跨平台编译,模板路径不正确
  - 使用 https://github.com/rakyll/statik 处理 js、css、image等静态资源
  - 使用 https://github.com/shiyanhui/hero 替换 html/template 模板引擎
- [x] redis查询问题
  - 缓存页面数据
- [x] 增加配置文件读取
  - 使用 https://github.com/spf13/viper
- [ ] Docker 部署
- [ ] goroutine 并发数控制
- [ ] 爬取数据的完整性


## Other
许多Go的原理还没弄懂，有精力会慢慢深究下。写得很潦草，多多包涵。
