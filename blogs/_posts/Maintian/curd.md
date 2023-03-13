---
title: crud
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---

# SQL boy

```sql
insert into test (account, password, username) values ('test', 'test',  'test')

delete from test where id>1

select account from test where account = ${data.account}

update test set username = 'test' where id = 2
```



##  insert

这个是插入返回的结果


```json
{
    "fieldCount": 0,
    "affectedRows": 1,
    "insertId": 10,
    "serverStatus": 2,
    "warningCount": 0,
    "message": "",
    "protocol41": true,
    "changedRows": 0
}
    
```

这里还会出现报错 当然了这里是因为格式不对

```cmd
{"code":"ER_TRUNCATED_WRONG_VALUE","errno":1292,"sqlMessage":"Incorrect datetime value: 'test' for column 'time' at row 1","sqlState":"22007","index":0,"sql":"insert into test (account, password, username, time) values ('test', 'test', 'test', 'test')"}
```

成功的信息

```cmd
{"fieldCount":0,"affectedRows":1,"insertId":11,"serverStatus":2,"warningCount":0,"message":"","protocol41":true,"changedRows":0}
```





## select





select 

是一个就是列表 返回全部的对象

```json
[
{
"id": 1,
"account": "dzg",
"password": "123456",
"username": null,
"avatar": null,
"signature": null,
"time": null
},
{
"id": 2,
"account": "admin",
"password": "123456",
"username": null,
"avatar": null,
"signature": null,
"time": null
}
]
```

## update

`let sql  = "update test set username = 'test' where id = 1"`

如果update是失效的 就也不会出现问题 这肯定是我们不想要的

通过检测是否有行被修改了 我们就可以判断是否是正确的sql语句了 

这里我觉得还是需要封装的

```json
{
"fieldCount": 0,
"affectedRows": 0,
"insertId": 0,
"serverStatus": 2,
"warningCount": 0,
"message": "(Rows matched: 0 Changed: 0 Warnings: 0",
"protocol41": true,
"changedRows": 0
}
{
"fieldCount": 0,
"affectedRows": 1,
"insertId": 0,
"serverStatus": 2,
"warningCount": 0,
"message": "(Rows matched: 1 Changed: 1 Warnings: 0",
"protocol41": true,
"changedRows": 1
}
```

## delete

```json
{"fieldCount":0,"affectedRows":5,"insertId":0,"serverStatus":2,"warningCount":0,"message":"","protocol41"
 :true,"changedRows":0}
```





## 原子操作

就是说 我们只操作我们的数据库 只会修改一行的数据





