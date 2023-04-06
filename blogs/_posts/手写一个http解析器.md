---
title: 手写一个http解析器
mathjax: false
date: 2023-04-03 21:03:14
tags:
categories:
---

# 手写一个http解析器

之前一直都是对于我们的http的格式一知半解,现在正好有机会可以学习java网络编程,就顺带着将其内容学完就好了.

我们先不考虑TCP三次挥手两次握手这样轱辘话我们就直接看socket.

# HTTP请求

## Request Header

```js
GET /js/app.dc20dbd2.js HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Host: localhost:8080
Sec-Fetch-Dest: script
Sec-Fetch-Mode: no-cors
Sec-Fetch-Site: same-origin
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36
sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: "Windows"
```


## Request Body 

```js
{"username":"123456",
"passwd":"123456"}
```


## Respond  Header

```js
HTTP/1.1 200 ok
Content-Type: text/html
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: *
Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept
```

## Respond Body

```js
{"filename":"package-lock.json","filedata":"data:application/json;base64,ewogICJuYW1lIjogIkFkbWluaXN0cmF0b3IiLAogICJsb2NrZmlsZVZlcnNpb24iOiAzLAogICJyZXF1aXJlcyI6IHRydWUsCiAgInBhY2thZ2VzIjoge30KfQo="}
```

如果用图像表示的话就是

`cr == \r    lf == \n`



![ ](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E9%94%AE%E5%85%A5%E7%BD%91%E5%9D%80%E8%BF%87%E7%A8%8B/4.jpg)

这样看来其实还是挺富有规则的感觉也不是很难啊.

# 一些关键术语

### Params

从?开始就是一个

```js
/api/download?filename=config.xml
/api/download/check?filename=name.txt&filadata=asdfasdfcfasf
```

这里我们先通过?解析出来,然后利用 & split 存到我们的hashmap利用 split = 然后分别交到里面去

通过我们的api里内部参数中,我只喜欢就是使用这个内容的功能

## Body

这里根据你的参数不同结果也不太一样.所以我们body的解析应该是放到最后去做的,也就是我们需要给我们的heandler 如何解析才可以知道

不然就是没有用的就是说 哈哈哈 

```js
{"filename":"package-lock.json","filedata":"data:application/json;base64,ewogICJuYW1lIjogIkFkbWluaXN0cmF0b3IiLAogICJsb2NrZmlsZVZlcnNpb24iOiAzLAogICJyZXF1aXJlcyI6IHRydWUsCiAgInBhY2thZ2VzIjoge30KfQo="}
```

```js
{"filename"
:"package-lock.json","filedata"
:"data:application/json;base64,ewogICJuYW1lIjogIkFkbWluaXN0cmF0b3IiLAogICJsb2NrZmlsZVZlcnNpb24iOiAzLAogICJyZXF1aXJlcyI6IHRydWUsCiAgInBhY2thZ2VzIjoge30KfQo="}
```

### AES

想cookie我们就是可以将其

