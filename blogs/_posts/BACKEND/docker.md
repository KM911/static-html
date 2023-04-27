---
title: docker
date: 2023-01-12 14:17:43
tags:
categories:
---


# 很有意思

image tag contain 

直接制作镜像

启动镜像 

```shell
docker run -p 3000:3000 my-docker-for-node
```



上传



将contain打包为镜像

```
docker commit CONTAINER ID <imagename>:<tags>
```



## 编写dockerfile

```dockerfile
# node 版本为12 的 -alpine 表示使用精简版
FROM node:12-alpine

WORKDIR /app

COPY . .
# RUN yarn install --production
CMD ["node", "/app/app.js"]
```

我们这里没有执行 `npm install` 是因为我们直接将 `node_moudles` 打包进去了 很好xiao



制作镜像

```cmd
docker build -t app:1.0 .  // 这个点很关键 笑死了
```

用自己的dockers







```
 netsh interface portproxy show v4tov4
```



# 使用docker 一键部署

分析一下 之前吧 我们使用`git` 对我们的项目进行部署 其实还是可以的 不是吗

