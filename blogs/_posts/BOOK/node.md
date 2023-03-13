---
title: nodejs
date: 2023-01-12 14:17:43
tags:
categories:
---

# 环境搭建

这里不解释了

# hello world

在我看来 这个才是真正意义上的hello world in nodejs –它解释了 我们的nodejs到底是做什么的 而不是你的npm环境 

创建一个server.js文件 写入下面的代码

```js
const http = require("http");

const server = http.createServer((request, respond) => {
    console.log("hello world");
    respond.end("hello world");
})
server.listen(3000);
```

在当前文件夹下打开cmd or powershell 输入 `node server.js`

然后打开浏览器 输入下面的网址 `localhost:3000/`

就可以看到 我们的cmd界面中 输出了`hello world` 浏览器显示了 hello world (不过是在屏幕的左上角 并且字体有些小)

## Nodejs

我们使用Nodejs 是开发网页的后端 (运行在服务器上的代码) 

用于处理我们浏览器的各种请求 比如http 是协议 get是协议中的method 吧

比如 http://zuoge.online 就是浏览器发送一个get请求到其对应的服务器上 服务器接收请求后 就会返回信息 显示在我们的浏览器中 可以是网页html 或者 数据 

## 解析request请求

我们的浏览器向服务器发送了http的请求 其中包含了很多信息 比如访问的具体网址 使用的设备 发送请求的ip 大体如图所示

```
```

其中最关键的是我们的路径解析 

## 路由

再看一段代码

```
const http = require("http");

const server = http.createServer((request, respond) => {
    console.log("hello world");
    respond.end("hello world");
})
server.listen(3000);
```



我们解释一下上面的代码

console.log() 在nodejs的运行窗口 输出

respond 返回的信息 就是浏览器访问我们的





## 获取get请求的数据



### 解析request中的数据

> 将req中的路由 解析 就是? 及其后面的内容
>
> > req.url  
>
> 将其解析为json
>
> > json_data = url.parse(req.url,true)
>
> 获取? 后面的参数
>
> > q = json_data.query
>
> 根据 键获取值
>
> > value  = q.key



## 获取post请求中的数据

post 无法从路径中获取参数

这里我们需要其他的方法

```js
// 先说js是如何发送post请求的
let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://localhost:3000/?action=add_article", true);
            xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
            xhr.send(JSON.stringify({
                "author": document.getElementById("author").value,
                "title": document.getElementById("title").value,
                "content": document.getElementById("content").value
            }));
            xhr.onreadystatechange = function () {
```

我们是如何接受post请求的





## 利用respond返回数据

### 设置状态信息

## 文件读取

fs.readFileSync

这是同步的读取 方式 避免发送空的信息回去

可以对普通文件进行发送 但是对于图片 该怎么办呢?



## 文件保存 



## 路径重新定向问题

有些路径的获取 会导致之前的路径无效

目前最好的解决方案就是用相对路径写 并且保证不要出现路径混乱 这里就是自己的代码写得烂啊 你自己想想啊 你如果重定向有问题 不就是直接gg了吗? 



> 重定向问题 
>
> > './重新定向文件夹/QQ图片20221004203229.jpg'

## 连接数据库

需要一个插件

```
npm install mysql --save 还是放到我们的工作环境里面吧
```

这里不写对于数据库的基本操作 而是写 对于操作的要点 比如说 如何去 



### 注册 和 登录

**注册方面**

* 有不同的账号 检查账号是否相同–账号已经被注册了
* 先不管这些吧 比如找回什么的



**登录**

user -> pwd -> yes 返回两个不同的信息

账号不存在 -> 密码错误 -> 登录成功

登录函数

```js
function login(respond, respond_data) {
    // 1. 获取到用户名和密码
    // 2. 判断用户名是否存在
    // 3. 如果存在就判断密码是否正确
    // 4. 如果密码正确就返回登录成功
    // 5. 如果密码错误就返回密码错误
    // 6. 如果用户名不存在就返回用户名不存在
    db.query("select * from user_info where username = ?", [respond_data.username], function (err, result) {
        if (err) {
            console.log(err)
        }
        else {
            if (result.length == 0) {
                // 说明没有相同的用户
                console.log("用户名不存在")
                respond.writeHeader(200, { "Content-Type": "application/json;charset=utf-8" })
                respond.end(JSON.stringify({ "code": 400, "msg": "用户名不存在" }))
            }
            else {
                // 说明有相同的用户
                if (result[0].password == respond_data.password) {
                    // 说明密码正确
                    console.log("登录成功")
                    respond.writeHeader(200, { "Content-Type": "application/json;charset=utf-8" })
                    respond.end(JSON.stringify({ "code": 200, "msg": "登录成功" }))
                }
                else {
                    // 说明密码错误
                    console.log("密码是 " + result[0].password)
                    console.log("错误的密码是 " + respond_data.password)
                    console.log("密码错误")
                    respond.writeHeader(200, { "Content-Type": "application/json;charset=utf-8" })
                    respond.end(JSON.stringify({ "code": 400, "msg": "密码错误" }))
                }
            }
        }
    })
}
```



## 解决跨域问题

我们需要在服务端 设置回复信息的头为

```
respond.setHeader("Access-Control-Allow-Origin", "*");
```

就算不设置 信息状态码 都是可以的



## 编写自己的package

这样可以化简我们的 

利用moudle.export{

函数名,}





# CMS项目

content management system 内容管理系统

我们的图床 百度网盘 其实都是一个CMS系统 

## NP 这里是图床 所以是公开的

分析我们的项目结构 

index 将上传图像 和 显示图像的页面 聚合到一起 

首先是一个 top 导航栏 好吧 

其实这里就已经🆗了 logo 是



我们需要一个log 分别是ico webp之类的地方 

| log显示部分    | 上传图片                                             |
| -------------- | ---------------------------------------------------- |
| 图床的显示部分 | 这里我们使用keep-alive 一个很长的时间 不会很快的更新 |

这里直接显示10张图片 显示文字 和 图片的内容

class = shower

如何对于图像 还是应该使用就是 new 2 old 的排列方式 图片的大小是统一的 这样让大家都可以看见 我们的缓存的信息 就是还是不错的 list



## 这里开始我们的上传头像文件作为试点

如何写一个合理的开发企划 

清晰明确的项目文档 哎

其实就是一个 caves 不是很难的吧 

所以我们看看 node 和 共

使用redis 进行优化 这里是 进行键值对的 存储 

不对 开始加载 然后就可以了 

消息队列 使用 redis 实现就好了这个其实估计没有什么技术难度 今年的寒假里 等等啊 我是不是就是经历过这个 这么又到 寒假了 其实我不是很喜欢这个东西 哎

  

## 其实我的问题有很多 第一是 

nodejs 真的不行吗? 我感觉还可以啊  

