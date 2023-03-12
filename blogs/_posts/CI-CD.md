---
title: CI/CD
mathjax: true
date: 2022-11-01 17:00:45
tags: 教程
cover: http://81.68.91.70/api/cover
categories: 
---

# CI/CD

新名词 但是 其实很好理解 比如说 我们的nodejs项目 最后的产品 可能是一个pkg打包完成的软件

我们将源代码push到仓库里 就会自动编译 打包 然后部署到生产环境 





从自己的GitHub小卡片到我们的第一个PR 到就是说 我们的vecelc 开始学习这些东西吧 好吧

先说一下 GitHub action 并不好用 为什么 因为 



## 静态页面部署

对于静态页面部署 可以说是没有任何的难度了 

只需要你创建仓库 然后导入就好了 像GitHub pages 和 vercel 都会自动的为你部署文件

## 动态页面部署

我们其实是可以利用api 实现动态页面的 但是有一些限制 就是你不能修改路由 

项目的结构 必须 和 网站的结构一致 （当然了 现在的我还没有接触过vue 和 react 这样的框架 很多话可能都是错误的）

## API 部署

首先我们要明白一件事情 我们的vercel.app是运行在一个只读的文件系统上的 所以我们不能进行存储的相关功能 但是其实还是的 可以保存到本地然后push上去就好了 (vercel好像有100GB的存储空间 真的用不完) 

Vercel 会自动创建一个server 我们只需要按照它的格式书写api就好了

`api/index.js`  

两种书写方式 ES6 和 commonjs

```js
// index.js commonjs
module.exports = async (req, res) => {
    // 其实这个也是ES6的语法 为什么会不兼容? 
  const { id, theme } = req.query; 
  const data = await getBilibiliInfo(id);
  data.theme = theme;
  res.setHeader('Content-Type', 'image/svg+xml');
  res.setHeader('Cache-Control', `public, max-age=${6000}`);
  return res.send(renderBilibiliCard(data));
};

```

为了实现本地的测试 我们可以创建一个 `app.js` 不然的话只能部署到vercel后才可以查看

```js
const express = require('express');
const app = express();

const index = require("./api/index")
app.use("/api",index) // get请求会交给我们的index 处理
app.listen(3000)
```

**ES6的写法** 

```js
// index.js ES6
export default async (req, res) => {
    res.end("hello world")
}
```

同样的我们创建一个`app.js`用于本地测试

```js
// app.js
import express from 'express';
let app = express();
import index from './api/index.js';
app.use("/api", index);
app.listen(3000);
```

引入其他模块

```js
// render.js 
function render(text) {
    return text + "调用了render方法";
}
export default render;
```

```js
// 在index.js 里写下
import render from "../render.js"
export default async (req, res) => {
    res.end(render("hello world"))
}
```

# 看看你的

这里我提供我的vercel api供大家使用

### GitHub 卡片

`http://learn-km911.vercel.app/api/?username=<GitHub用户名>`

![](http://learn-km911.vercel.app/api/?username=km911&theme=react&show_icons=true)

### 常用语言卡片

`https://learn-km911.vercel.app/api/top-langs/?username=<用户名>&layout=compact`

![](https://learn-km911.vercel.app/api/top-langs/?username=KM911&layout=compact&theme=vue)



# 项目参考

[GitHub Readme Stats](https://github.com/anuraghazra/github-readme-stats)
