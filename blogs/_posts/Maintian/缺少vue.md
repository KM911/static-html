---
title: vue项目的常见报错汇总
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---



## 缺少`@vitejs/plugin-vue`

出现了上面的原因其实很简单

我们的vue其实是不能运行的 需要将其打包变成普通的`html`文件才可以被浏览器正常解析 

所以vue提供了`vite`这个工具给我们打包 与此同时`vite`还具有预览的能力 自己提供后端 所以我们需要写一个

`vite-config.js`文件

```js
import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'


export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})

```

