---
title: mysql环境搭建
date: 2023-03-12 14:17:43
tags:
categories:
---


# MySql 环境搭建

# 环境搭建

我不想花费很多的精力在这种事情上了 这是运维的事情

我们应该将重点放在数据库的设计上 这里需要我们立即业务



## 开启远程连接

首先 我们的mysql密码是在哪里的 

本来是在安装的时候设置的 但是在服务器端是一个什么情况 这里不理解 最好笑的是我不能本地登录 只能远程登录了 哈哈哈哈

```sql
mysql -u root -p
use mysql;
update user set host = '%' where user = 'root';

```

这里结果是 不能自己在本地 进入 了 我笑死了 哈哈哈



我们是用ip连接的 只需要输入ip 就好了





## 无法用密码登录




```js
ALTER USER 'root'@'localhost' IDENTIFIED BY '@Dzg15484' PASSWORD EXPIRE NEVER; #修改加密规则
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '@Dzg15484'; #修改密码
FLUSH PRIVILEGES; #刷新数据
```



