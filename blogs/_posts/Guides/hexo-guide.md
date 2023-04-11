---
title: hexo_guide

date: 2023-03-12 12:06:05
tags: 教程
categories: 
---

# hexo 教程

时隔半年,我又想起了hexo[^为什么回来了] 

因为过去了半年,自己对于很多事情的看法也发生了变化,比如我觉得博客的核心是里面的内容,而不是封面的妹子有多好看:joy: .其实也是收到了比如 [阮一峰]() 等人的影响,他们的博客可能并不好看,甚至可以说是丑陋,但是里面的内容深度真的让人拍手叫绝,所以我也开始将精力放到内容上了.



## 环境配置

hexo是基于nodejs的静态博客生成工具,所以你需要安装nodejs,并将其添加到你的环境变量中.



## 初始化

```
hexo init
```

## 下载主题

```
git clone -b master https://github.com/jerryc127/hexo-theme-butterfly.git themes/butterfly
```

## 配置主题

在_config.yml中修改 theme

```
theme: butterfly
```

修改基本信息

```
title: 标题
subtitle: '副标题'
description: '描述'
keywords: '关键词'
author: 作者
language: zh-CN
timezone: ''
```

修改主页链接 就是GitHub pages的url

```
url: https://km911.github.io/BLOG
```

配置md的模板 (scaffolds/post.md)

```
---
title: {{ title }}
date: {{ date }}
 #开启数学公式
tags:  # 标签
cover:  # 缩略图
categories:  # 分类
---
```



## 主题butterfly下的配置

头像

```
```

follow me的GitHub主页

```
```





## 吐槽

真的很想吐血 有的时候 我们的path是没有用的

是不是就搞我一下 人糊涂了 都要





## 主题优化

其实butterfly 已经为我们配置好了很多的优化选项 我们只需要去下载插件 然后开启就好了 

### 汇总

* 本地搜索
* 点击特效 
* 数学公式
* 文字统计
* 背景效果
* 音乐播放



### 本地搜索

### 点击特效

### 数学公式

### 文字统计

### 背景效果

### 音乐播放

播放音乐的原理 其实是访问那些音乐提供商 比如网易之类的 我们主要是利用这个可以给我们的网站加速 我觉得 毕竟这个就是用你自己的网络去访问了 而不是去访问我们的服务器 我们的服务器只给了url 

```html
<div class="aplayer no-destroy" data-id="7248481267" data-server="netease" data-type="playlist" data-fixed="true" data-autoplay="true" data-lrctype="0"> </div>
```

想这个就是我们一个模板

[参考教程](https://butterfly.js.org/posts/507c070f/#%E6%8F%92%E5%85%A5-Aplayer-html) 

我们现在写一个专门播放单曲的

````
<div class="aplayer no-destroy" data-id="761319" data-server="netease" data-type="song" data-fixed="true" data-autoplay="true" data-lrctype="0"> </div>
````

<div class="aplayer no-destroy" data-id="761319" data-server="netease" data-type="song" data-fixed="true" data-autoplay="true" data-lrctype="0"  > </div>

[^为什么回来了]: 其实是想用hugo的,但是hugo对于新人不是很友好,就回来了,哈哈哈哈:joy:
