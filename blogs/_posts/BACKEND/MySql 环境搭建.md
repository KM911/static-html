---
title: mysql环境搭建
date: 2023-03-12 14:17:43
tags:
categories:
---


# MySql 环境搭建

# 环境搭建

我不想花费很多的精力在这种事情上了 这是运维的事情 我们的重点应该是对于数据库的性能负责就是.

### 创建用户并给予权限

`host`的值为 % 表示允许全部的ip访问

```sql
create user 'username'@'host' identified by 'password'; 
# 直接给全部的权限 这样是最简单的不是吗
grant all privileges on *.* to 'username'@'%' with grant option; 
```

## 无法用密码登录


```js
ALTER USER 'root'@'localhost' IDENTIFIED BY '@Dzg15484' PASSWORD EXPIRE NEVER; #修改加密规则
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '@Dzg15484'; #修改密码
FLUSH PRIVILEGES; #刷新数据
```

drop user 'pjuser'@'host' 

## 开启远程连接

首先 我们的mysql密码是在哪里的 

本来是在安装的时候设置的 但是在服务器端是一个什么情况 这里不理解 最好笑的是我不能本地登录 只能远程登录了 哈哈哈哈

```sql
mysql -u root -p
use mysql;
update user set host = '%' where user = 'root';
```

这里结果是 不能自己在本地 进入 了 我笑死了 哈哈哈

我们是用ip连接的 只需要输入ip 就好了 或者你直接允许全部的ip访问 ,不过考虑到我们需要使用代理,可能导致我们无法正常登录

redis创建副本

## Mysql 版本问题

mysql 6.0 以后驱动版本直接变成了 
