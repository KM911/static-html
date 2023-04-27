---
title: RPC比HTTP要慢?
mathjax: false
date: 2023-04-27 17:34:17
tags:
categories:
---



# RPC VS HTTP go语言实现



## 理论分析

`RPC` 和 `HTTP` 都是基于 `TCP`协议的,所以两者在 *网络链路层*  *网络层* *传输层* 都是一样的,唯一不同的地方就是 *应用层*的协议不同

两者共同的开销 : 以太网头部（14个字节） + IP头部（20个字节或更多）+ TCP头部（20个字节或更多） 

` HTTP`协议的开销 : 主要是集中在请求头,当然了,出于效率的考虑,你可以减少携带的请求头数量. (50个字节或者更多)

```
POST /hello HTTP1.1
Content-type:text/html;
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36
....
Body
```

`RPC`协议的开销: 函数名,参数.几乎没有其他额外的开销.

```
type Call struct {
	ServiceMethod string     // The name of the service and method to call.
	Args          any        // The argument to the function (*struct).
	Reply         any        // The reply from the function (*struct).
	Error         error      // After completion, the error status.
	Done          chan *Call // Receives *Call when Go is complete.
}
```

可以发现 `RPC`的开销是更小的,但是也不多,我个人认为两者的性能差距是很小的,不过我的测试结果似乎不太相同.

## 设计实验

简单来说就是分别开启 `HTTP`服务器 和 `RPC`服务器,利用go语言自带的测试框架比较其性能就好了.

代码发布在[GIthub](https://github.com/KM911/Benchmark)和[Gitee](https://gitee.com/dong-zuoge/Benchmark),欢迎pr.



### HTTP服务器

```go
func HttpServer() {
	// 创建一个http服务器
	server := http.Server{
		Addr: config.HttpServerPort,
	}
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		length := request.Form.Get("length")
		atoi, _ := strconv.Atoi(length)
		writer.Write([]byte(CreateString(atoi)))
	})
	server.ListenAndServe()
}
```

### 发送HTTP请求

```go
func HttpRequest() {
	url := "http://localhost" + config.HttpServerPort + "/hello?length=10"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 这里的开销也有一部分 我们可以将其删除 然后比较其结果 发现影响不大 只有百分之几的结果影响
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Host", "localhost:5000")
	//req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(body)[:0])
}
```

### RPC服务器

其实从代码量上看,我觉得RPC比HTTP要简洁很多,似乎性能应该很好吧?!

```go
func RPCServer() {
	rpc.RegisterName("API", new(RPCAPI))
	listener, err := net.Listen("tcp", config.RPCServerPort)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	rpc.Accept(listener)
}

type RPCAPI struct {
}

func (r *RPCAPI) Hello(length int, reply *string) error {
	*reply = Http.CreateString(length)
	return nil
}
```

### 发送RPC请求

```go
func RPCRequest() {
	client, err := rpc.Dial("tcp", "localhost"+config.RPCServerPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()
	var reply string
	err = client.Call("API.Hello", 10, &reply)
	if err != nil {
		log.Fatal(err)
	}
}
```

### 其他说明

我这里还定义了一个函数 `CreateString` ,作用是和 `sleep` 类似的,增加单次请求响应的时间,让结果差距变得更加明显.

```go
func CreateString(length int) string {
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		builder.WriteString("a")
	}
	return builder.String()
}
```

### Benchmark

```go
func BenchmarkRunHttp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Http.HttpRequest()
	}
}
```

```
go test -bench=. -benchmem  ./test
```

结果很令我出乎意料,`RPC` 比 `HTTP` 慢了 太多了,rpc执行一次的时间几乎是http的8倍.不仅如此,rpc还使用了更多的内存.

```
goos: windows
goarch: amd64
pkg: httpvsrpc/test
cpu: AMD Ryzen 5 5600H with Radeon Graphics
BenchmarkRunHttp-12        21718             54578 ns/op            3599 B/op         46 allocs/op
BenchmarkRunRPC-12          2979            391075 ns/op           20576 B/op        277 allocs/op
PASS
ok      httpvsrpc/test  2.977s
```

## 分析原因

首先,从内存分配次数上看,我们的RPC就要比我们的http分配次数多4倍.我们找找看,RPC是在哪里分配了内存.

### 内存分配

在我们发起RPC请求之前,我们需要先创建一个 `client` ,然后通过 `client` 去发送 `RPC`请求

```go
type Client struct {
 codec ClientCodec

 reqMutex sync.Mutex // protects following
 request  Request

 mutex    sync.Mutex // protects following
 seq      uint64
 pending  map[uint64]*Call
 closing  bool // user has called Close
 shutdown bool // server has told us to stop
}

```

并且创建 `client`的这个过程中,是需要去上锁,进一步带来了额外的开销


```go
func (client *Client) input() {
	var err error
	var response Response
	for err == nil {
		response = Response{}
		err = client.codec.ReadResponseHeader(&response)
		if err != nil {
			break
		}
		seq := response.Seq
		client.mutex.Lock()
		call := client.pending[seq]
		delete(client.pending, seq)
		client.mutex.Unlock()
        ...
```

也就是说RPC的开销很多是来自创建我们的 client那我们现在修改一下我们的测试用例

我们应该发送多次请求,这样RPC就可以复用它的 client ,我们现在改成单个测试用例发送100次请求

```go
	var reply string
	for i := 0; i < 100; i++ {
		err = client.Call("API.Hello", 10, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}
}
```

```go
for i := 0; i < 100; i++ {

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Print(string(body)[:0])
}
```

再跑一次我们测试用例,结果和我们地预期差不多,我们地RPC实现了反超,并且分配了更少地内存,看来我们一开始的猜想是对的,我们RPC请求中创建client会花费大量的时间和内存.

```
BenchmarkRunHttp-12          222           5365716 ns/op          308948 B/op       4215 allocs/op
BenchmarkRunRPC-12           295           4052054 ns/op           41952 B/op        772 allocs/op
```

我们将请求的次数放大到1000次,看看效果.差距被进一步拉开,两者占用的内存内存相差15倍.

```
BenchmarkRunHttp-12           19          54527200 ns/op         3134236 B/op      42721 allocs/op
BenchmarkRunRPC-12            32          35607600 ns/op          236561 B/op       5272 allocs/op
```

## 总结

RPC确实是一种比HTTP更为轻量化的应用层协议,特别是在大量的请求中,两者的差距将会更加明显





欢迎大家进行交流,共同提高.
