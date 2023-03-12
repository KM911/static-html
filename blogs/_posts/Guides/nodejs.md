---
title: nodejs环境搭建
mathjax: true
date: 2022-09-15 09:56:46
tags:
cover: http://81.68.91.70/api/cover
categories: 教程 
---



## nodejs环境搭建

[下载node](https://pc.qq.com/search.html#!keyword=node) 然后安装 请牢记你的软件安装路径 很关键

## 配置环境变量

添加nodejs的环境变量 具体参考我的[如何添加环境变量](https://km911.github.io/BLOG/2022/09/15/环境变量)

你需要做些什么? 

1. node
2. node_modul
3. npm

打开cmd 输入

```
node -v
```

 如果出现了版本信息 说明配置成功



## npm 换源

我们使用`npm` 下载的地址 

```js
npm get registry // 查看当前使用的地址
npm config set registry XXX //修改为XXX
npm config set registry https://registry.npm.taobao.org  //使用淘宝的地址
```



