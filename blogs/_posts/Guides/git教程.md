---
title: git教程
date: 2022-09-11 09:06:16
tags: 博客
categories: 教程
---

# 环境搭建



## git下载和安装

这里推荐去腾讯下载 毕竟国内的网速相对较快

[腾讯软件下载](https://pc.qq.com/search.html#!keyword=git)

下载和电脑匹配的位数 32位和64位 (记得点普通下载 而不是高速下载)

[64位](https://webcdn.m.qq.com/spcmgr/download/Git-2.37.3-64-bit.exe)

[32位](https://webcdn.m.qq.com/spcmgr/download/Git-2.37.3-32-bit.exe)

![image-20220911091043294](http://localhost:3000/picgo/1668762156326.webp)

如果打开控制台 输入git 会有反应 说明配置成功了



## 基本信息配制

这里不要用我的信息 啊 (其实就算用了也不会出什么问题)

```
git config --global user.name "KM911"
git config --global user.email 2547715095@qq.com
```

### 关闭服务器代理 

很关键 不然容易443 time out 不过其实就算用了 还是会是不是443的 没办法 谁让人家微软不是中国企业呢 w

```
git config --global --unset http.proxy
git config --global --unset https.proxy
```

### 生成密匙

```
ssh-keygen -t rsa -C 2547715095@qq.com
ssh-keygen -t rsa -C '2547715095@qq.com' -f D:\0.SOFT\ssh\.ssh\github
ssh-keygen -t rsa -C '2547715095@qq.com' -f D:\0.SOFT\ssh\.ssh\gitee

这里其实可以设置ssh的生成位置 默认c盘 就算没了 再重新生成一个就好了
```

#### 关于生成密钥的过程中(可选)

我们可以指定生成的密钥位置(很关键 可以移动到D盘)和名字(不关键)

在上述指令后 输入

```
D:/0.SOFT/ssh/.ssh/key
```



然后我们需要去设置git 去告诉它我们的密匙的位置

```
git config --global core.sshCommand "D:\0.SOFT\ssh\.ssh\"
ssh-add D:\0.SOFT\ssh\.ssh\gitee
这里是因为我的key就是叫做key

我想到了 其实我们可以控制变量
1: 不设置我们的key 只是修改目录 (共用一个key 不修改邮箱 ) 
2: 设置不同的邮箱 生成两个key 

```



## 添加密钥



---



### 关闭ssl认证

其实还是为了解决网络 问题 最好还是科学上网

这里推荐一个(不是广告 我自己也在用)

[pigcha]() 按流量付费的话 30元 15g 我用了好久 还是因为拿去看V才用光的 啊哈哈哈 如果只是用来上github 还有一个软件可以使用

[steam++]() 好像改名字了 不过这个不重要 免费软件 可以给开发者一点支持 让他买杯咖啡 不过因为是DNS代理或者修改host 容易出现浏览器无法使用的情况 对于新人来说 问题有点多 当然了 steam++还是很强大的 

```
git config --global http.sslverify false
```



### 基本使用流程

先在GitHub上创建一个仓库 然后将其clone到本地

```
git clone 仓库的链接
```

将想要上传的文件添加到缓存区 提交到本地后 上传到仓库

```
git add filename
git commit -m "updata"
git push 
```

稍等一下 你就可以在你的仓库 看见你的文件了 

反正我就是只是使用这些最基本的命令 依然很强大 分支什么的对于我们来说还没有必要 

恭喜你 你基本上已经学会了git的80% 可以愉快的玩耍了 

 

## 其他参考教程



安装git 这些东西都不是问题啊 其实我还是推荐 你去使用我们ubuntu系统来学习 因为有 apt get 这种很方便的命令行指令 (windos也有 不过支持相对没有那么好 毕竟术业有专攻)

[参考教程_雪峰](https://www.liaoxuefeng.com/wiki/896043488029600)

[参考教程_知乎_](https://zhuanlan.zhihu.com/p/30044692)



# 情景再现

[参考教程](https://www.jianshu.com/p/5f99380597de)

在具体的情况下学习 有助于我们理解和记忆具体的内容

## 获取云端新的分支

`git branch -a` 查看本地和云端的分支 

`git branch -b new_baanch` 切换到一个新的分支  

`git fetch --all` 获取云端的全部修改

`git branch --set-upstream-to=origin/new_branch` 将该分支进行追踪新的分支

`git pull` 获取新的内容

## 本地和云端的代码冲突问题

直接放弃本地的方法

```
git fetch --all  
git reset --hard origin/master
```

当然了 这样肯定不是很好 我们应该可以这样解决



## 缺少文件没有提交

将本次提交 和上一次提交合并 

但是如果你已经将代码提交到云端了 就不要在修改了 

```
git commit --amend
```



git rebase



git cherry-pick



# 工作流



为了保证多人合作项目的时候 不会影响master分支的代码 我们使用如下的工作流 

下面的指令 需要在**当前文件夹**下 打开cmd powershell 或者git-bash 进行输入

## 克隆仓库到本地

```
git clone https://gitee.com/dong-zuoge/node.git
```

## 开辟自己的分支

```
git checkout -b <your-branch>
```

![image-20221014211305053](https://i0.hdslb.com/bfs/album/2a9a4da5fda911f3430143a8d753d517d2d7045e.png)

这样就算你修改了文件 不会影响主分支

## 查看自己修改

你可能修改了一些代码 但是自己记不清了  使用如下的指令进行查看修改

```
git diff
```

我们可以借助一些可视化的工具[Github Desktop](https://desktop.github.com/)

如果只是修改了几行 或许还可以使用原生的git 但是如果对于项目进行了大量的修改 并且涉及到不同层级的文件时 还是比较推荐使用这种工具

## 将修改提交到本地

将修改提交到本地 这里我们还是没有修改云端 需要等待项目的主要负责人将其合并

```
git add *
git commit -m "提交的信息 最好可以描述清楚 自己做了什么修改 如何实现的 当然了其实就算很简单 也没有关系"
```

![image-20221014220330151](http://localhost:3000/picgo/1668762175877.webp)

左上角的对应我们的 git add * 

GitHub Desktop 默认会将全部的修改进行提交  你也可以选择不提交某些文件[^gitignore]

左下角对应我们的 git commit -m “提交的信息”

这里就是让我们输入提交信息的地方 下面的是Description(描述) 更加详细的描述

[^gitignore]: 我们可以专门创建一个`.gitingore`文件 告知什么文件不要进行提交

## 将提交合并(不重要)

有的时候 我们只是做了一点点小的修改 如果还是进行一样的提交方式 就会产生很多无效的commit

为了简化我们的提交信息 我们可以这样做 

```
git add * 
git commit --amend
```

这样就会将本次的提交合并到上一次提交上 

这里会进入我们的编辑模式 让你查看修改 

`CTRL`+`C` 进入命令模式 然后输入 `qa` 就会退出了 

不要在已经将commit 上传了以后进行此操作 会出现比较复杂的修改

尽管不会影响项目的运行 但是会产生垃圾信息

![image-20221014222626888](https://i0.hdslb.com/bfs/album/03d12b99fb8418ed5ba0f55469867e1a221620f9.png)

## 将修改提交到云端

如果你以及确保自己的修改已经完毕了 (这期间里你可能进行了多次的 add commit的操作)

![image-20221014214558020](https://i0.hdslb.com/bfs/album/2628cec54613bc4caf8ed00d477bcbb7f74f2cf4.png)

将修改提交到云端

```
git push 
```

对应的操作是 点击紫色的按钮 

这里特别将 上面的`Pull origin`按钮圈出来 是想说明 这里两个箭头的含义 向上的箭头表示 我们本地有多少次提交没有上传到云端 同理 向下的箭头 表示云端有多少修改没有同步到本地![image-20221014221814260](https://i0.hdslb.com/bfs/album/853cb81c06843cdcc3679919ba2b3ece79755cb8.png)



## 将分支合并(不重要)

这部分内容其实很简单的 因为都有可视化

进入我们的[仓库地址](https://gitee.com/dong-zuoge/node)

![image-20221014223615538](https://i0.hdslb.com/bfs/album/8806e9b053d999a2fffe3269b1773acafde481d4.png)

`Pull Request` 申请将分支合并 

这里需要等待项目的负责人和审查员进行检查后 才会进行合并 

为了防止我这样的菜鸟直接修改master分支 影响系统的运行 很有必要

![image-20221014225035523](https://i0.hdslb.com/bfs/album/f927dc8b71aff07f8e098325f921476c57bf00a3.png)

这里推荐使用 扁平化合并 

这样不会将其他分支的每一次提交信息都合并到master分支上 

可以保持我们的master分支的整洁性

## 移除本地的分支

等到我们发现 自己的pull request 被合并以后 (说明你的修改是没有问题的) 我们也不再需要该分支了

进行如下的操作  将本地分支移除 

我们的仓库就变成了 只有master分支的样子了 

```
git checkout master 
git branch -D <your-branch>
```

# 总结

其实对于一般的情况 就比如说 我们个人开发 我们只需要进行

```
git add *
git commit -m "提交的信息"
```

多次提交和修改之后 

```
git push 
```

完成了 是不是很简单啊 但是我还是推荐你养成良好的git使用习惯 这样对于自己代码回滚 还是为工作打下良好的基础 都是很好的练习 

# 题外话

工具是为了提高效率 我并不觉得使用gui就一定 怎么这么样 使用命令行就 巴拉巴拉

 如果你觉得使用某种方式更加适合自己 有助于提高自己的效率 这样就够了 没有必要花费时间和别人争论怎么做更好之类的



就算你对于一个工具不是非常了解 就比如这里的git 和 Github Desktop 

掌握一些基本的技巧 就可以体验到新的工具 带来的好处 所以多看看 有什么新奇的玩意 然后学习一下吧

 



