---
title: sql锁
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---

# 并发中的mysql

因为并发出现的问题 

脏读 不可重复读 幻读



事务的隔离级别

## 读锁和写锁

读锁 共享锁: 只能读 , 不可写

写锁 互斥锁: 不能读 , 只有当前事务可以写