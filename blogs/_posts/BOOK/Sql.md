---
title: SQL
date: 2023-01-12 14:17:43
tags:
categories:
---

## SQL

我们不是第一次接触了 所以我想以一种面向使用场景的方式来进行记录

你其实不需要知道什么是 DDL DML这些 也不需要知道如何create a database 这些并不能帮助你了解和认识我们数据库

这里我们推荐你使用navicat进行数据库可视化的创建和管理 在vscode中是使用sql notebook编写sql语句

## 后端查询方式

* 普通查询 select
* 条件查询 where
* 模糊查询 like
* 排序 order by desc 
* 范围查询
* 分组查询 
* 分页查询
* 链接查询
* 合并查询
* 子查询



## 登录和注册

我们规定接口格式如下

```
/login/?username=admin&password=123456
/register/?username=admin&password=123456
```

首先我们应该明白,对于像密码格式检查比如只允许字母+数字 不能存在特殊字符 像这种应该在客户端实现 这样的好处在于可以更加快速检测错误 还可以减轻服务端的压力.

包括像对于password之类的重要信息进行加密 都可以将其放到客户端 这里应该使用非对称的加密算法 这样即使数据库的数据被他人盗窃,用户的密码也不会丢失.



### 表的设计

我们可以很自然的设计出这样的表格式.

```sql
CREATE TABLE `users`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER  NOT NULL,
  `password` varchar(32) CHARACTER  NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
)
```



### 注册功能的实现

```go
// 这里部分的知识是web开发的也就顺带着写一下吧
func Register(c *gin.Context) {
	// 获取query的参数
	username := c.Query("username")
	password := c.Query("password")

	// 查询数据库中是否存在该用户
	var user model.User
	dao.DB.Where("username = ?", username).First(&user)
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"code":     "1",
			"message":  "该用户名已被使用",
		})
		return
	} else {
		dao.DB.Create(&model.User{Username: username, Password: password})
		c.JSON(200, gin.H{
			"code":    "0",
			"message":  "注册成功",
		})
		return
	}
}

```

这里的逻辑是很简单的 先利用username查询数据库中是否存在数据 如果没有则插入数据 注册成功 反之返回注册失败 该用户名已被使用

### 登录

一开始我是这样写的

```go
func Login(c *gin.Context) {
	// 获取query的参数
	username := c.Query("username")
	password := c.Query("password")
	// 查询数据库中是否存在该用户
	var user model.User
	dao.DB.Where("username = ? and password = ?", username, password).First(&user) 
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "登录成功",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":    "1",
			"message": "用户名或密码错误",
		})
		return
	}
}

```



```go
func Login(c *gin.Context) {
	// 获取query的参数
	username := c.Query("username")
	password := c.Query("password")
	// 查询数据库中是否存在该用户
	var user model.User
	dao.DB.Where("username = ?", username).First(&user)
	if user.Id != 0 {
		if user.Password == password {
			c.JSON(200, gin.H{
				"code":    "0",
				"message": "登录成功",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code":    "1",
				"message": "密码错误",
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"code":    "1",
			"message": "用户不存在",
		})
		return
	}
}
我的其实和它的是一样的,也是只有三个可能,只不过它的返回的是 用户名或者密码错误,这里也确实是有这个可能的,但是为了避免混淆还是直接说明就是密码错误就好了
```

我个人觉得后者更好,首先两种方法都只查询了一次数据库,性能上相差不大.后者有更加详细的错误提示,可以引导用户进行正确的操作.

第一种用户就无法知道自己是密码还是账号错误了,就不知道是该检查账号是否正确还是选择忘记密码.



## 视频列表

视频列表格式如下 

```sql
CREATE TABLE `videos`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `likes` int(0) UNSIGNED  NOT NULL DEFAULT 0,
  `comments` int(0) UNSIGNED NOT NULL DEFAULT 0,
  `url` varchar(255)  NOT NULL,
  `title` varchar(255) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
)
```

我们的视频数据有很多可能是上万条或者更多,但是用户肯定是不需要这么多的,我们每次只发送5条数据给用户即可 

我这里是根据id来进行排序 你也可以添加比如时间字段 然后按照时间的先后顺序进行发送

```go
dao.DB.Where("id > ?", last_time_id).Limit(5).Find(&videos)
```



## 用户视频列表

这里很明显是一个多表结构了 准确讲是一对多,可能一个用户有多条视频,也可能一条都没有

这里我们需要重新定义我们的表结构 当然了 其实只需要修改视频表就够了

```sql
CREATE TABLE `videos`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0,
  `likes` int(0) UNSIGNED  NOT NULL DEFAULT 0,
  `comments` int(0) UNSIGNED NOT NULL DEFAULT 0,
  `url` varchar(255)  NOT NULL,
  `title` varchar(255) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
)
```

其实也只是多加了一个user_id字段,但是就是这一个字段,就可以让用户表和视频表产生关联

比如投稿接口/publish?userid&title 实现如下

```go
func Publish(c *gin.Context) {
	//获取参数
	title := c.Query("title")
	userid := c.Query("userid")

	//插入数据库
	dao.DB.Table("videos").Create(&model.Video{
		Title:    title,
		UserID:   userid,
		URL:      "http://www.baidu.com",
		Likes:    0,
		Comments: 0,
	})
	c.JSON(200, gin.H{
		"code":    "0",
		"message": "发布成功",
	})
}
```



## 评论列表

我们的数据库的cell是平行的关系 但是我们的评论是嵌套关系(因为存在回复该条评论的评论 不过一般不会出现更深层级的评论了)

你可能会说 用一个字符列表去存储评论不就好了 但是你要知道在关系型数据库中 每个单元的长度是固定的 如果你的二级评论很多 就会出现问题 

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment_list": [
        {
            "id": 0,
            "user": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true
            },
            "content": "string",
            "create_date": "string"
        }
    ]
}

```

// TODO 二级评论的实现 这个其实也不困难 只要不是实现回复了谁 或者自己显示它的id 将这个 其实目前这个功能还是问题比较多的





# 索引

我们其实一直都在使用索引,比如使用id作为主键 就是一个索引

比如说登录接口我们首先需要去检查用户名是否存在 这个时候无法使用id去检查,如果你未username建立索引,可以去利用username快速获取到

## 分类

主键索引 id 只能有一个

唯一索引 username 

常规索引

全文索引 



聚集索引 : 必须有,只能有一个 就是将我们的索引和数据放到一块 主键就是聚集索引 如果不存在主键,就会利用第一个唯一索引作为(id)

二级索引 : 只有索引的结构(其实还携带了id)通过id进行索引 效率还是很高的

### 例子

回表查询

```sql
select * where username=123456;
```

首先通过username获取到id再通过id去获取全部的信息.





### 获取数据库的操作次数

```sql
show global status like 'Com_______';
```

返回的结果为

![image-20230221155332736](http://81.68.91.70/tinypicgo/image/1676966012.webp)

很明显 可以看出来 select是使用最多的 其余还有 commit 

### 慢查询日志

mysql会自动记录执行时间过长的记录

默认关闭  开启方式 修改

![image-20230221155637882](http://81.68.91.70/tinypicgo/image/1676966198.webp)

### profile

![image-20230221175150167](http://81.68.91.70/tinypicgo/image/1676973110.webp)

## 最左前缀法则 

我们建立的复合索引 最左边的值必须存在





这里我还没有理解复合索引 

先试试看 就是 非唯一的索引 建立了索引

其实就存储多个id 然后再去查 也可以实现增加速度的



### 索引失效的原因

*   没有使用单引号 username = 123456 和 username = ‘123456’ 是不一样的
*   使用了运算符 使用了> 可以使用= 或<=



前缀索引 对于大量文本的内容 我们可以通过就是只获取部分前缀的方法建立索引





![](http://81.68.91.70/tinypicgo/image/1676982899.webp)





索引的选取 你建立索引更多 未必就是一定就好的 你知道





## 4.2 功能测试

在本次项目的功能测试中，我们主要针对项目所涉及的各项功能进行了全面测试，包括但不限于以下几个方面：

1.	用户注册与登录,我们对用户注册与登录功能进行了测试，确保用户能够正常注册账号，登录账号，并能够进入系统。我们还针对一些异常情况进行了测试，例如输入错误的用户名或密码、尝试使用已存在的用户名注册等。

2.   视频上传和视频流,我们测试了用户上传视频后是否可以生成视频封面,视频流能否正常播放,非登录用户也可以获取视频流,同时无法进行其他需要权限的操作.

3.   视频点赞和评论,包括用户能否正常对视频进行点赞和评论,视频点赞数,评论数是否进行了及时的更新.

4.   关注用户和聊天,我们对关注用户和发送消息进行了测试，包括用户能否正确接受和发送信息、历史信息能否正常显示。我们还测试了在发送消息过程中可能发生的异常情况，例如已经将用户取消关注等。

通过以上功能测试，我们发现该系统的各项功能均能够正常工作，符合用户需求，并且能够在各种异常情况下正确处理。

## 4.2.1测试用例

1.  用户注册测试用例：

-   输入正确的用户名和密码，注册成功。
-   输入已经存在的用户名，注册失败并提示用户名已经被使用。
-   输入无效的用户名或密码，注册失败并提示错误信息。
-   在注册过程中，中途退出并重新开始，注册流程能够正确恢复。

2.   用户登录测试用例：

-   输入正确的用户名和密码，登录成功并跳转到用户详情页。
-   输入错误的用户名或密码，登录失败并提示错误信息。
-   在登录过程中，中途退出并重新开始，登录流程能够正确恢复。
-   在已登录状态下，直接打开系统主页，能够正确进入系统而不需要再次登录。

3.   视频流测试用例:

*   没有登录时,可以正常获取视频流,同时无法其他需要权限的工作.
*   已登录,在获取视频流的前提下,可以正常使用其他功能,例如对视频点赞等.

4.   其他接口测试用例

*   未登录,无法通过JWT的认证,会提示错误信息.
*   已登录,可以在正常响应并返回正确的信息.

## 4.3 性能测试

在本次项目的性能测试中，我们主要针对系统的性能进行了测试，包括但不限于以下几个方面：

1.并发性能测试 我们对系统进行了并发性能测试，通过模拟多个用户同时访问系统，测试系统是否能够正常处理并发请求，以及并发请求数量的上限。测试结果表明，该系统在并发请求量较小时能够正常工作，但在并发请求量较大时会出现卡顿或者响应时间过长的情况。

2.响应时间测试 我们对系统的响应时间进行了测试，包括用户登录、任务发布和接受、任务完成等操作的响应时间。测试结果表明，系统响应时间较为稳定，但在高并发情况下响应时间会变长。

3.资源占用测试 我们测试了系统在不同负载下的资源占用情况，包括CPU、内存和网络带宽等。测试结果表明，系统在高并发情况下会占用较高的CPU和内存资源，但对网络带宽的占用较小。

根据以上性能测试结果，我们建议在系统的性能优化方面，可以采取以下措施：

1.增加系统的并发处理能力，提高系统的吞吐量，减少卡顿和响应时间过长的情况。 2.对系统的代码进行优化，减少不必要的资源占用，提高系统的性能表. 3.针对数据库的查询进行进一步优化,提高系统的响应速度.

项目的结构

![image-20230222104735879](http://81.68.91.70/tinypicgo/image/1677034056.webp)





# 多线程

在go中你可以很轻松地开始协程 利用

```go
go func() // 这里就不会阻塞了

```

## 互斥锁

互斥锁可以有效解决多线程对于资源修改的问题,比如去银行取钱,只能运行一个线程在对其进行操作.

两行代码就可以解决这个问题.

```go
Lock.Lock()
defer Lock.Unlock()

```

互斥的问题在于,无法保证go routing的顺序,我们的程序被block后,执行的结果是不确定的

## 通道

这里的还有一个问题 我们想让我们的程序按照一定的顺序去执行多线程任务 这里其实不对吧 , 对的 你可以开两个



## 异步

异步的前提就是多线程,如果没有多线程,就会直接卡死了,这里其实是开启了一个隐藏的线程,待其完成后,再进行对应的操作.

```go
func test(x, y int, callback func(a, b int) int) {
	println(callback(x, y))
}

// 回调函数的具体实现
func add(x, y int) int {
	return x + y
}

func main() {
	x, y := 1, 2
	test(x, y, add)
}

```

简单就是我们将一个函数名作为参数,再在该函数内部调用,并不是真正意义上的回调,因为你可以比如说在一开始就调用该函数,但是如果我们直到最后才调用,就是一个回调,并且拿到返回值比我们的nodejs要简单太多了 不是吗.

所以我们可以直接 这样就保证了回调函数的正确性.

```go
defer callback()
```

用a用户注册 存在相同的用户 注册成功 随机的账号

用b用户登录 用户名不存在 然后密码错误 最后登录成功   1111 托尼有点忧伤

最后登录c用户查看消息发送. 123456789 一算 

我绝对欧克







## SQL预编译

其实我就一直都在使用只是我自己没有意识到,笑死我了.

```java
// 使用PreparedStatement实现SQL预编译
String sql = "SELECT * FROM user WHERE id=?";
PreparedStatement pstmt = conn.prepareStatement(sql);
pstmt.setInt(1, 1); // 给参数设置值
ResultSet rs = pstmt.executeQuery();
while (rs.next()) {
    // 处理结果集
}

// 使用Statement实现普通SQL语句执行
String sql = "SELECT * FROM user WHERE id=1";
Statement stmt = conn.createStatement();
ResultSet rs = stmt.executeQuery(sql);
while (rs.next()) {
    // 处理结果集
}

```

当然了很多SQL框架都可以进行SQL预编译的操作,并且将其简化了

```java
// mybitas中使用SQL预编译的例子
@Update("update goods set status = 1 where id = #{id}")
public void DropGoods(int id);
```

SQL预编译的好处如下：

1. 提高性能：SQL预编译可以将SQL语句预处理成一个可重复利用的模板，当需要执行相同的SQL语句时，只需要传入不同的参数即可，避免了每次执行都需要解析和优化SQL语句的开销，提高了执行效率。
2. 防止SQL注入：使用SQL预编译时，所有的参数都会被当作参数占位符传递给SQL语句，而不是将参数直接拼接在SQL语句中，避免了SQL注入攻击。
3. 代码可读性：使用SQL预编译可以让代码更加清晰易懂，通过将SQL语句和参数分离，可以使代码更易于维护和调试。

SQL预编译的缺点如下：

1. 内存占用：在执行预编译语句时，需要为每个预编译语句分配内存，如果预编译的语句过多，会占用大量的内存。
2. 网络开销：由于预编译语句的结果需要传输到客户端，所以在网络传输方面会增加一些开销，特别是当预编译的语句较多时，会占用更多的网络带宽。

总的来说，SQL预编译是一种提高数据库操作性能的有效方式，通过将SQL语句预处理成一个可重复利用的模板，可以避免每次执行都需要解析和优化SQL语句的开销，提高了执行效率，并且可以防止SQL注入攻击。但是，预编译语句的过多会占用大量的内存，并且增加网络开销。









