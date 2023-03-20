---
title: ffmpeg简单使用
mathjax: false
date: 2023-03-18 19:17:17
tags:
categories:
---

# 提取视频封面

```go

```

# 压缩视频

```shell
ffmpeg -i input.mp4 -c:v libx265 -preset medium -crf 28 -r 24 output.mp4

```

```shell
ffmpeg -i %1 -c:v libx265 -preset medium -crf 28 -r 24 H265%1
```

