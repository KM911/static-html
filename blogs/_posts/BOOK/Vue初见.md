---
title: vue初见
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---

# 环境搭建

两种方式

1 CDN引入

考虑到 CND引入 每一次都需要下载 我们可以直接将该js文件下载到本地使用

```html
 <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
```

2本地化的Vue

```js
npm install --save vue
```

vue和react ts之类的`语言`一样 是无法被浏览器直接运行的 我们需要将其编译成为浏览器可以直接运行的 原生`js`

## hello world

不想多写 可是不得不写

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
    </script>
</head>

<body>
    <!-- 但是这个页面并不能显示出vue的特性 响应式数据绑定 -->
    <div id="app">
        <h1>{{ message }}</h1>
    </div>
    <!-- <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script> -->
    <script src="vue.js"></script>
    <script>
        var app = new Vue({
            el: '#app',
            data: {
                message: 'Hello Vue!'
            }
        })
    </script>

</body>

</html>
```

在我看来 一个更加好的 hello world

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script>
    </script>
</head>

<body>
    <!-- 这个页面就很好的体现了vue的特性 我们只是修改了count的值 就实现了页面的更新 在之前 我们需要手动的去更新页面 -->

    <div id="app" @click="count">
        <h1>被点击了 {{nums }}次</h1>
    </div>
    <script src="vue.js"></script>
    <script>
        var app = new Vue({
            el: '#app',
            data: {
                nums:0
            },
            methods:{
                // 控制num进行自增
                count:function(){
                    this.nums++;
                }
            }
        })
    </script>

</body>

</html>
```

## 数据更新 到 视图更新

v-html v-text v-bind  v-moudel v-once

```html
```

 双向数据



视图的显示

v-if v-show



v-on 绑定函数

 

v-for 生成大量数据 很完美





## 案例

一个存储在服务器上的todo list

首先是前端页面的设计 我们需要一个比较美丽的效果

