---
title: nginx
date: 2023-03-12 14:17:43
tags:
categories:
---

1.  关于静态文件托管的问题

nginx中可以利用 location +root 实现静态文件服务 web框架中也提供了静态文件服务static 请求采用哪一种更好呢?  

2.   nginx中有

# Nginx

我们先不用管它是什么 我们先看看我们如何使用它 我们现在有了一台服务器 假设它的ip是 192.168.1.5 现在你用浏览器去访问它 你会发现得不到任何的响应

很好理解 你的服务器并不知道是否需要回应的你的消息以及如何回应 



现在我们开始进行配置nginx     对nginx conf 进行配置

## hello world

```nginx
worker_processes 1;
events {
    worker_connections 1024;}
http{
  include mime.types;
  default_type application/octet-stream;
  sendfile on;
  keepalive_timeout 65;
  server {
    listen 80;
    server_name localhost;
    location / {
      root /www/wwwroot/app/public/BLOG ;
      index index.html index.htm;
    }
  }
}
```

## location

假设我们的 css文件都是 web/css/*.css 这样的形式存储的 我们可以写这样的语句 原因和简单 因为我们的 这里它还是会将全部的 url作为参数 而不是匹配后的

```
location /css{
	root /www/web
}
```

  

让我们去看看吧 它存放在我们的nginx目录中的

```nginx
types {
    text/html                                        html htm shtml;
    text/css                                         css;
    text/xml                                         xml;
    image/gif                                        gif;
    image/jpeg                                       jpeg jpg;
    application/javascript                           js;
    application/atom+xml                             atom;
    application/rss+xml                              rss;

    text/mathml                                      mml;
    text/plain                                       txt;
    text/vnd.sun.j2me.app-descriptor                 jad;
    text/vnd.wap.wml                                 wml;
    text/x-component                                 htc;
    ......
```

很明显 这里是为了说明文件的类型  错 而是为了决定如何处理我们的文件 是对该文件进行下载 还是运行js 或者渲染html

错误处理 

## sendfile

是否进行缓存 还是直接发送 这里很明显对于大批量请求 还是需要做缓存的 小请求可以就是

## GZip Br

```nginx
  gzip on;
        gzip_min_length  1k;
        gzip_buffers     4 16k;
        gzip_http_version 1.1;
        gzip_comp_level 5;
        gzip_types     text/plain application/javascript application/x-javascript text/javascript text/css application/xml;
        gzip_vary on;
        gzip_proxied   expired no-cache no-store private auth;
        gzip_disable   "MSIE [1-6]\.";
```

## Tengin concat

将我们的静态文件请求合并 这里其实也好理解

因为我们的每一个http请求都需要创建会话 这里就