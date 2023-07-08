---
title: TODO
top: true
mathjax: false
date: 2023-04-03 13:07:34
tags:
categories:
---

#  Stream

-   [ ] headers 为代理服务 
-   [ ] 添加更多的指纹信息
-   [ ] 性能优化 这里其实我觉得就是减少部分的异步服务







## 利用nginx解析gzip请求

我们知道对于静态资源,我们利用Gzip进行压缩,可以减少网络资源的消耗(当然了,这会增加部分CPU资源的消耗) 

同样的,我们可以将我们的请求进行压缩,特别是对于`JSON` 格式,理论上会有非常不错的效果.

我们可以通过一个实例看看 测试的结果是对于小文件,例如4kb以内的,压缩效果并不是很理想 只有请求大于4kb时使用比较好

我们现在有两个解决方案 第一个是利用例如springboot或者是java内置的ungzip方法,对其进行解压缩

第二个方法是利用nginx 这样就不会因为编程语言而被限制了 

问题解决了 其实非常简单 根本就不需要什么lua脚本 其他的UnGzip 只需要在你的nginx下添加一句

```nginx
   location / {                 
    gunzip on;                                     
	proxy_pass http://localhost:3000;                                                                                    }     
```

通过抓包 方向分析 内容 将我们的proxy 修改为我们自己的 理论上来说 发送http请求 根本就不想需要

还有这里有upd浏览 他们那里 什么 TLS WS 都是什么乱七八糟的协议啊 就是 其实感觉没有很大的必要



1. 日志问题: 日志分类 怎么写 什么时候写 

   日志大致分类 info error waring 具体什么情况下才是 因为我感觉只有info 和 error 

   怎么写: id + data 但是这里的id其实是data中的字段,类似于订单的id,这个id会在流程的接口中不停流转 

   什么时候写: 一般是访问 + 运行结果



1. 性能优化: 同步 日志 网络 预热 

同步: 线程创建 

日志:如果可以减少部分无用日志

网络:因为需要更新js 这里会导致单词请求花费时间过长

预热: JIT 类似于jvm的

多进程: 因为nodejs的主线程是单线程的 无法充分利用服务器性能 可以利用child_pross模块去提升性能 但是资源占用会高很多

```
{"url":"https://wsdkdl.migu.cn:8443/rocher/83c8a141d7064b5b9f93e3b7a5715a56.js?v=2V2bfQ-Oh4JazIJsfhxpOr",
"headers":[{"key":"X-Proxy","value":"192.168.0.120"},{"key":"Real-ip","value":"localhost"}],
"__s":"25cfca035d0f47828a1edbab5ef34dfc","ID":"202306281903242810","type":"MUSIC","canvas":"iVBORw0KGgoAAAANSUhEUgAAADIAAAAeCAYAAABuUU38AAACrklEQVRYR+2Yv08UQRTHP8/EwtBoghBzRhuJGq00VirGisKc2tsZyKnRVhqxQBtpNXKEE/8FgdCaqLX+AUonJOo29JiMeXvLMXu3ezuzc3uhcKq7nfd9731258ebEYMxdLW/wDtgA/gGbGt/20p/6qONiGPfR9m5BlwG7nT7SP6vAl+BTyJ8zrEZyGNpg0jHWQt4Bvy23Y8A94ClH8Ay8J5xIl4C0+5pRLEQlkXYdJe5WaZAngBvsnSngSkF6e18DLx2i2VbLQLzIvzyl2YrOiC5EKq7DpzPBtHukjAqnRaJR3Fwi0FaCDN5ro4kw+pQPohKdcB5DLPU1xHhUSiJ7GLMSSQ9J2yvk8C55EHG0NozHQd+AofLZbQuwu1y0rZKmhjzwJrsKWcXgKvWkz4gatUEGuWzCYKROsasZ4F0Q2iCBSB1YK08iCoXyw4zqWHMtg2ic+KKNZzsxApAasBWGIiqZ0TQXcCrCbqPjAiMAqeAs4BO7KxWAKKSnt3VK52O8QnfpVkwxtDY3xD7xh0eiPcQU5At2KzBRLl3Z6kGNLT2PE74VAAKsgazdVgIBhnAZLdzWBBh1jUpBWnAWBO0FAprgctvd/BIhDHXjCQyR29eZOdjqkh0VVt2gRtiXsRJEb64pCPGMNeC+dwSxcVLWInSL8JzEV64pKAgH/Q80bdoLPAUUDQW5bgqwt0iI+1XkM7SXwamQog4f5G8+imNlwLRrsyDVcYr0TnhebByebE9NqVB1NMusNJ91AV0n7gE3ALul690vYCCQLwiVWz8H6TiF+zt3ueLxMuvd4ThCLyW3zm90RhOXt5RvDZEvSOp9PLMO/19wQ3Xi734IGIMf4DjAQGrkPoVjQnIK+BpFdkE+PQs49tf5Ayg96EHqXkerJLUjeEt8PCAkHgfdf8Bm3S1HzV+Bf4AAAAASUVORK5CYII=","userAgent":"Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-A6060) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/15.0 Chrome/90.0.4430.210 Mobile Safari/537.36","language":"zh-CN","hardwareConcurrency":"8","deviceMemory":"-1","platform":"Linux aarch64","doNotTrack":"unknown","colorDepth":"24","screenWidth":"412","screenHeight":"892","screenAvailWidth":"412","screenAvailHeight":"892","devicePixelRatio":"2.625","indexedDB":"1","openDatabase":"1","maxTouchPoints":"5"}
```

## 代理服务

因为 不同的使用方法是不一样的 尽管他们都可以称自己为代理服务 因为确实也满足了 这个要求 我现在其实就是想利用抓包 去看看 具体的内容


# Task

-   [x] 有自己一套键盘规则
  
-   [ ] VSCode插件 中文符号替换
  
-   [x] 一个本地化的文件共享工具 
  
-   [ ] 网站防抖 
  
-   [ ] GO语言的ORC
  
-   [ ] 利用我们之前的天气卡片
  
-   [x] Github action 自动化部署我们的hexo
  
-   [x] 动态的封面图 (但是存在BUG 短时间内的大量 请求会导致我们的封面变成同一张图片)

-   [x] LPA Local Page Application 

一种借助我们的浏览器进行开发的软件 

好处和坏处我们自己去分析
-   [ ] github release 
  
-   [ ] 用go重写我们的图床软件 但是并不是很好用的呢
  
-   [x] 利用go写cli工具 进行vue项目的初始化的创建 




# Fin

-   [x] 学习使用hugo 改为使用hexo好吧 并且学会了使用GitHub action

-   [x] 视频编码问题以及解决了 无法使用GPU只能CPU硬解码

-   [x] 将博客的文章全部整理好 整理了 但是没有整理好 不是吗 哈哈哈哈 就是有点僵

-   [x] go 发布自己的package 主要是为了做备份 其余都不重要 哈哈哈

-   [x] 命令行工具有了 









