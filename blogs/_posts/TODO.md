---
title: TODO
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
top: true
tags:
categories:
---



#  stream



-   [ ] 学习使用hugo 改为使用hexo好吧
-   [x] 视频编码问题以及解决了 无法使用GPU只能CPU硬解码





















# Task

-   [x] tiktok极简版
-   [ ] school 这个也是可以的 是吧 
-   [ ] java 这个可以先敷衍一下 其实主要是 我不是很喜欢java 不过可以将其做为真正的项目开始接触
-   [ ] 有自己一套键盘规则



首先 这里的IO操作是异步的,json会直接返回,如果这个时候的图片没有加载出来 就会报错了

我们的ffmpeg是对本地的视频文件做处理,这个肯定是需要等待的



这里客户端是不是进行了处理 如果投稿的时间过长 就直接报网络的错误 



我们可以去延后 投稿列表接口



## todo



-   [ ] VSCode插件 中文符号替换

-   [ ] 一个本地化的文件共享工具 

-   [ ] GO语言的ORC

-   [ ] 利用我们之前的天气卡片

-   [ ] Github action 自动化部署我们的hexo

-   [ ] 动态的封面图 (但是存在BUG 短时间内的大量 请求会导致我们的封面变成同一张图片)

-   [ ] LPA Local Page Application 

    一种借助我们的浏览器进行开发的软件 

    好处和坏处我们自己去分析

-   [ ] github release 

-   [ ] GitHub app and so

-   [x] 开始使用我们的faker alice 了

-   [x] 用go重写我们的图床软件 但是并不是很好用的呢

-   [x] 就是成为一个键盘侠 将自己的全部软件进行一个用命令行进行输入然后打开 基本上是完成了

-   [x] 利用go写cli工具 进行vue项目的初始化的创建 

  

# 键盘侠

我们现在知道了 如何利用go语言去执行cmd的指令

学会了复制文件路径 和文件夹

现在利用start . 可以打开当前目录就是说

我们现在开始维护一张映射表吧

start ink 就可以打开当前的软件了

```markdown
g Google
f fileManger 
fl 文件的路径 打开相应的路径 其实这里我觉得就是说 不是很方便的结果
w wechat
t tim
s shortcut file 就是直接打开我们的
d D盘 
c C盘


这样其实很麻烦的结果其实是次要的结果 有这么多的程序其实我们还是需要

```

## 实现方式

我们维护了一个字典 进行一个映射 就可以直接打开了

我们这里还有功能没有实现其实 比如快速文件夹



从自行车的修车 对于问题的思考



我其实之前自己就是修过自行车 就是自己瞎琢磨 也得到了一些宝贵的经验 但是不是很关键 并不能很好得帮助我们解决核心问题 我去B站上找了一个视频 其实很简单 就花费了大概10分钟的时间 我就完全明白该如何做了 ‘

其实我们自己去写代码也是一样的 自己写固然很好 但是你还是需要先看过别人写一次 你才可以知道该如何做 懂了吗 

现在快快 开始向其他人进行模仿吧 加油 我相信你 的 就是最强的



# 中文符号替换

这里我们直接遍历全局就好了

一键将中文的符号换成英文 这个是我的痛点 不知道我们是不是可以写出来 

用js 

目前已有的插件是针对实时输入的 无法解决我们的 问题 等等 它不会检测就是复制的内容 所以 我:smile:

*   检测复制 将复制的内容修改
*   一键替换 将全局的内容进行替换



# github action

```
name: GitHub Actions Demo


on: push


jobs:
  my_first_job:
    name: My first job
  my_second_job:
    name: My second job
```



# Whether-card

技术分解: 爬虫+SVG动画+Vercel

- [ ] 获取天气信息 axios
- [ ] SVG动画 PS or
- [x] Vercel 托管我们的api

最晚在我们的寒假写出来吧

# 基于GO语言的本地ORC

发现现在的ORC基本上都需要上传至服务器 我想实现一个本地的

# 动态的封面图

我想到了一个就是api获取随机的图片url这样会不会可以呢?试试看

OK 完成了 就是花费了我们一个两个小时的时间 其实我觉得还可以 

文件型数据库 