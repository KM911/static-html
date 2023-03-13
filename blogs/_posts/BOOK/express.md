---
title: express
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---

# express

一个nodejs的后端框架 但是目前 我并没有体会到 它的好用之处

`hello world`

```js
const express = require('express');
const app = express();

app.get('/', (req, res) => {
    res.send('Hello World!');
}
);
app.listen(3000, () => {
    console.log('Example app listening on port 3000!');
}
);

```

然后是它很奇怪的路径传值的方式 我们一般是利用 url +?action= …

express是直接使用url的全不 好像不支持? 的方法 感觉有些呆? 

占位符传参 就是说 /直接包裹信息/ 这个不太好理解就是接口是做什么的 

## 等我们的项目足够大了 或许我会再用这个重新写一次

## 参考文档

[Express官方文档]()

## 吐槽

可能是我太菜了 对于前端的理解过于浅显了 在学习使用`express`的时候总感觉 教程是给哪些已经学会了人使用的 对于我们这样的普通人 阅读体验实在是不是很好

hello world是很简单的 但是很快就会让你摸不着头脑了 哎 还是自己的水平有限吧

# 环境搭建

我不推荐新人就直接使用脚手架之类的东西 我们从零开始 一点点学习和理解 不然真的看不懂的

```shell
// 创建一个文件夹 并下载express
mkdir express 
cd express
npm install --save express
```

# hello world in express

创建一个名为hello.js文件

```js
const express = require('express');
const app = express();
app.get("/", (request, response) => {
    response.send("hello world");
});
app.listen(3000);
```

```powershell
node hello.js
```

然后用浏览器打开我们的 `localhost:3000`

就可以看见hello world了

到这里的教程还是很正常的 我们下一步是学会托管静态页面

## 托管静态文件

创建一个public 文件夹 在里面写一个名为index.html的文件 

```js
app.use(express.static('public'));
```

这样我们就可以直接访问`public`下的静态文件了

`localhost:3000/index.html` 不必再考虑相对路径了

## 如何实现访问限制

就是说 我们只想`public `下的`css`文件夹被访问

```js
app.use("/css",express.static('public/css'))
```

我们访问css的请求还是不会出现问题 同时也对访问权限进行了设置

# 处理请求

```js
app.get("/:action/:name/:pwd"){
    //这样不就可以了 其实也差不多的 哎
}
```

##
