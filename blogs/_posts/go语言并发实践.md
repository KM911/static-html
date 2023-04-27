---
title: go语言并发实践
mathjax: false
date: 2023-04-27 17:33:30
tags:
categories:
---

# go并发问题实验

**之前做项目的时候第一次遇到了并发问题,但是没有解决,现在就来弥补一下当时的遗憾吧.**

**本人不过是一个刚刚接触go的新人,如果存在错误,请及时指出,不胜感激.**

# 实验设计

## 实验环境

* **win 10**
* **go 1.19 **
* **mysql 8.0**
* **fiber **
* **gorm**

**实验代码可以访问 **[gitee]()

## 问题场景

**用户可以对一个视频点赞,假如现在有两个用户同时点赞.**

```
A:
count := sql.query("like")  // 获取当前的点赞数 假如为 10 
count ++ // 11
sql.update(count)// 更新为 11
B:
count := sql.query("like")  // 因为是同时(或者是在更新之前)进行了查询 count为10 
count ++ // 11
sql.update(count)// 结果为11 很明显不正确 
```

## 代码结构

* **config 监听的端口 mysql的dsn **
* **controller 处理请求**
* **dao 和数据库交互**
* **test 测试脚本 用于发送http请求.**

## 实验配置说明

**DSN是一定需要修改的,主要是填写mysql的用户名和密码.**

**mysql默认端口号是3306,请检查一下.**

```
KM911LocalDSN = "root:@Dzg15484@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

KM911LocalDSN = "root:@Dzg15484@tcp(localhost:3306)/go?
"<username>:<password>@tcp(localhost:3306)
```

**然后就是创建数据库**

```
# 登录数据库
mysql -u root -p 

create database go;

```

**可以直接导入demo.sql文件,或者是复制下面的内容**

```
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for demo
-- ----------------------------
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `like` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
INSERT INTO `demo` VALUES (1, 0);
INSERT INTO `demo` VALUES (2, 0);
INSERT INTO `demo` VALUES (3, 0);

SET FOREIGN_KEY_CHECKS = 1;
```

## 并发安全测试

**主要是涉及到关于视频点赞数更新的操作.**

**下面是三中不同的更新数据库中数据的方法.**

### 查询后更新

```
func UpdateLikeAfterQuery() {
var like LIKE
DB.Table("demo").Where("id = 1").First(&like)
like.Like = like.Like + 1
DB.Table("demo").Where("id = 1").Update("like", like.Like)
}
```

### 查询时更新

```
func UpdateLikeWithoutQuery() {
DB.Exec("update demo set `like` = `like` + 1 where id = 2")
}
```

### 互斥锁

```
func MutexLike() {
MutexLock.Lock()
var like LIKE
DB.Table("demo").Where("id = 3").First(&like).Update("like", like.Like+1)
MutexLock.Unlock()
}
```

### 实验结果

```
go run test/test.go  
```

> **查询后更新点赞数**
> **用户id: 1 查询点赞数**
> **   结果为 12**
> **用户id: 1 点赞数+100,查询点赞数**
> **   结果为 19**
> **查询时更新点赞数**
> **用户id: 2 查询点赞数**
> **   结果为 8137**
> **用户id: 2 点赞数+100,查询点赞数**
> **   结果为 8237**
> **使用互斥锁**
> **用户id: 3 查询点赞数**
> **   结果为 4850**
> **用户id: 3 点赞数+100,查询点赞数**
> **   结果为 4950**

**可以发现,在查询时更新数据和使用互斥锁都可以保证并发的安全,下面对其二者进行性能测试.**

## 性能测试

**根据上面测试的结论我们已经知道如何保证并发安全了,只需要计算一段时间内,点赞数增加量的多少就可以知道哪一个方法的性能更好了.(类似相同时间内,跑的距离更长的人跑得更快)**

**这里可能会因为每个人的电脑性能不同,比如你的电脑性能很好,2秒就将2000个并发任务全部完成了,就无法体现出两者的性能关系,可以加大循环次数或者减少等待时间.**

```
func BenchMark(id string) {
println("用户id:", id, "查询点赞数")
Query(id)
for i := 0; i < 2000; i++ {
go Like(id)
}
time.Sleep(time.Second * 2)
Query(id)
}
```

### 测试结果

**查询时更新点赞数变化为分别为1337 1777,因为两者相差比较大,肯定不是测试误差,我后续还测试了几次,使用互斥锁的性能要好20%左右. **

> **在查询时更新点赞数**
> **用户id: 2 查询点赞数**
> **        结果为 0**
> **        结果为 1337**
> **使用互斥锁**
> **用户id: 3 查询点赞数**
> **        结果为 0**
> **        结果为 1777**

# 实验总结

**在对数据库进行更新操作时,需要特别注意并发问题.**

**使用互斥锁和查询时更新数据都可以保证并发安全,同时互斥锁的性能要稍好.**
