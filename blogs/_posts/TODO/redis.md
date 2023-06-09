---
title: redis
mathjax: false
date: 2023-04-28 18:31:04
tags:
categories:
---

## redis安装

像这样的常用服务端软件,一般的包管理工具都会提供.其实不必太担心.

```
yum -y install redis
```

## 启动redis

```shell
redis-server config
```

这里我们编写一个就是systemctl

```
[Unit]
Description=Redis
After=network.target

[Service]
Type=forking
ExecStart=/usr/local/bin/redis-server /etc/redis/redis.conf
ExecReload=/usr/local/bin/redis-server -s reload
ExecStop=/usr/local/bin/redis-server -s stop
PrivateTmp=true

[Install]
WantedBy=multi-user.target
                                       
```

```
[vagrant@xf-dev ~]sudo vim /usr/lib/systemd/system/redis.service
[Unit]
Description=The redis-server Process Manager
After=syslog.target network.target

[Service]
Type=forking
PIDFile=/var/run/redis_6379.pid
ExecStart=/usr/local/redis/bin/redis-server /usr/local/redis/etc/redis.conf
ExecReload=/bin/kill -USR2 $MAINPID
ExecStop=/bin/kill -SIGINT $MAINPID

[Install]
WantedBy=multi-user.target
```



## 最简配置

```
requirepass 123456
bind 0.0.0.0
daemonize yes
databases 1

dbfilename database.rdb
dir /root/redis/data
rdbcompression yes
logfile /root/redis/log
```





![image-20230430084417189](http://81.68.91.70/pg/image/KMdv8ylTK5nA.png)

![image-20230430084515975](http://81.68.91.70/pg/image/KMnGu2yK32ug.png)







![image-20230430085041078](http://81.68.91.70/pg/image/KMBDHSqdwGr8.png)









![image-20230430085600325](http://81.68.91.70/pg/image/KMdWSQiYWoNw.png)

![image-20230430085936608](http://81.68.91.70/pg/image/KMcBZXFIkV1R.png)

![image-20230430090234537](http://81.68.91.70/pg/image/KM8j0kUVibuX.png)







![image-20230430091130264](http://81.68.91.70/pg/image/KMP1D9mlw11E.png)







### redis 持久化

redis将自己的数据复制一份保存到磁盘中,下次启动会自动加载 

- ### RDB 

- ### AOF持久化

- 混合持久化

平常使用AOF进行持久化 





### 内存回收机制  内存淘汰机制

当我们的内存使用到达上限了,我们就需要进行回收.这里是决定如何进行回收的地方.

* LUR Least Used Resently
* LFU  Least Frequent Used



![image-20230430094151976](http://81.68.91.70/pg/image/KMFuvZRwoNKj.webp)

![image-20230430094012193](http://81.68.91.70/pg/image/KMx0R6VrF4mZ.png)



