# 服务端

```js
/ 获取到cookie 
let cookie=req.headers.cookie
/ 设置cookie
res.cookie = "key=value" || res.cookie = "key1="
```



# 客户端

问题还有 我们的cookie 是利用什么格式存储的 

所以 我们不能存储中文 和 特殊符号 这里需要进行一个转换

```js
/ 为了保护信息 你只能读 不能写
let cookie = document.cookie
/ 并且 这时的cookie 是一个字符串 形式为
''
```

