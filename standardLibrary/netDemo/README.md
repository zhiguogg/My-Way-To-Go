# net/http

Go语言内置的`net/http`包十分的优秀，提供了HTTP客户端和服务端的实现。

## HTTP协议

超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议，所有的WWW文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。

## HTTP客户端

### 基本的HTTP/HTTPS请求

Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。

```go
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

程序在使用完response后必须关闭回复的主体。

程序在使用完response后必须关闭回复的主体。

```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```



### GET请求示例



```
func Get(url string) (resp *Response, err error)
```

Get向指定的URL发出一个GET请求，如果回应的状态码如下，Get会在调用c.CheckRedirect后执行**重定向**：

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
```

如果c.CheckRedirect执行失败或存在HTTP协议错误时，本方法将返回该错误；**如果回应的状态码不是2xx，本方法并不会返回错误。如果返回值err为nil，resp.Body总是非nil的**，调用者应该在读取完resp.Body后关闭它。



使用`net/http`包编写一个简单的发送HTTP请求的Client端，代码如下：

```go
func getDemo()  {
	resp,err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	header := resp.Header
	for k,v := range header{
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))
}
```

我们的浏览器其实就是一个发送和接收HTTP协议数据的客户端，我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，然后浏览器会按照HTML、CSS等规则将网页渲染展示出来。


### 带参数的GET请求示例

关于GET请求的参数需要使用Go语言内置的`net/url`这个标准库来处理。

```go
func main() {
	apiUrl := "http://127.0.0.1:9090/get"
	// URL param
	data := url.Values{}
	data.Set("pageNo", "1")
	data.Set("pageSize", "500")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
```



### net/url

```go
//ParseRequestURI函数解析rawurl为一个URL结构体，本函数会假设rawurl是在一个HTTP请求里，因此会假设该参数是一个绝对URL或者绝对路径，并会假设该URL没有#fragment后缀。（网页浏览器会在去掉该后缀后才将网址发送到网页服务器）
func ParseRequestURI(rawurl string) (*URL, error) {}

type URL struct {
	Scheme      string
	Opaque      string     // 编码后的不透明数据
	User        *Userinfo // 用户名和密码信息
	Host        string    // host或host:port
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // 编码后的查询字符串，没有'?'
	Fragment    string    // 引用的片段（文档位置），没有'#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```

URL类型代表一个解析后的URL（或者说，一个URL参照）。URL基本格式如下：

```
scheme://[userinfo@]host/path[?query][#fragment]
```

scheme后不是冒号加双斜线的URL被解释为如下格式：

```
scheme:opaque[?query][#fragment]
```

注意路径字段是以解码后的格式保存的，如/%47%6f%2f会变成/Go/。这导致我们无法确定Path字段中的斜线是来自原始URL还是解码前的%2f。除非一个客户端必须使用其他程序/函数来解析原始URL或者重构原始URL，这个区别并不重要。此时，HTTP服务端可以查询req.RequestURI，而HTTP客户端可以使用URL{Host: "example.com", Opaque: "//example.com/Go%2f"}代替{Host: "example.com", Path: "/Go/"}。



```go
type Values map[string][]string
```

Values将建映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，**Values的键是大小写敏感的**。

```go
func (v Values) Encode() string
```

Encode方法将v编码为url编码格式("bar=baz&foo=quux")，**编码时会以键进行排序**。

```go
	v := url.Values{}
	v.Set("source","suggest")
	v.Set("q","进击的巨人")
	s := v.Encode()
	fmt.Println(v)
	fmt.Println(s)

// output
map[q:[进击的巨人] source:[suggest]]
q=%E8%BF%9B%E5%87%BB%E7%9A%84%E5%B7%A8%E4%BA%BA&source=suggest
```



```go
func (u *URL) String() string
```

String将URL重构为一个合法URL字符串。

```go
//现在想要访问 https://www.douban.com/search?source=suggest&q=进击的巨人

	apiUrl := "https://www.douban.com/search"
	uri, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println(err)
	}
	v := url.Values{}
	v.Set("source","suggest")
	v.Set("q","进击的巨人")
	s := v.Encode()
	uri.RawQuery = s
	get, err := http.Get(uri.String())

```



### Post请求示例

```go
// Post向指定的URL发出一个POST请求。bodyType为POST数据的类型， body为POST数据，作为请求的主体。如果参数body实现了io.Closer接口，它会在发送请求后被关闭。调用者有责任在读取完返回值resp的主体后关闭它。
func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
   ...
}
```



```go
func postDemo()  {
   // url
   url := "http://127.0.0.1:9090/post"
   // 表单数据
   //contentType := "application/x-www-form-urlencoded"
   //data := "name=小王子&age=18"
   // json
   contentType := "application/json"
   data := `{"name":"小王子","age":18}`
   resp, err := http.Post(url, contentType, strings.NewReader(data))
   if err != nil {
      fmt.Printf("post failed, err:%v\n", err)
      return
   }
   defer resp.Body.Close() //关闭
   b, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Printf("get resp failed, err:%v\n", err)
      return
   }
   fmt.Println(string(b))
}
```



### 自定义Client

要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：

```go
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://example.com")
// ...
req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```



```go
type Client struct {
    // Transport指定执行独立、单次HTTP请求的机制。
    // 如果Transport为nil，则使用DefaultTransport。
    Transport RoundTripper
    // CheckRedirect指定处理重定向的策略。
    // 如果CheckRedirect不为nil，客户端会在执行重定向之前调用本函数字段。
    // 参数req和via是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）。
    // 如果CheckRedirect返回一个错误，本类型的Get方法不会发送请求req，
    // 而是返回之前得到的最后一个回复和该错误。（包装进url.Error类型里）
    //
    // 如果CheckRedirect为nil，会采用默认策略：连续10次请求后停止。
    CheckRedirect func(req *Request, via []*Request) error
    // Jar指定cookie管理器。
    // 如果Jar为nil，请求中不会发送cookie，回复中的cookie会被忽略。
    Jar CookieJar
    // Timeout指定本类型的值执行请求的时间限制。
    // 该超时限制包括连接时间、重定向和读取回复主体的时间。
    // 计时器会在Head、Get、Post或Do方法返回后继续运作并在超时后中断回复主体的读取。
    //
    // Timeout为零值表示不设置超时。
    //
    // Client实例的Transport字段必须支持CancelRequest方法，
    // 否则Client会在试图用Head、Get、Post或Do方法执行请求时返回错误。
    // 本类型的Transport字段默认值（DefaultTransport）支持CancelRequest方法。
    Timeout time.Duration
}
```

Client类型代表HTTP客户端。它的零值（DefaultClient）是一个可用的使用DefaultTransport的客户端。

Client的Transport字段一般会含有内部状态（缓存TCP连接），因此Client类型值应尽量被重用而不是每次需要都创建新的。**Client类型值可以安全的被多个go程同时使用。**

Client类型的层次比RoundTripper接口（如Transport）高，还会管理HTTP的cookie和重定向等细节。



```go
func (c *Client) Do(req *Request) (resp *Response, err error)
```

Do方法发送请求，返回HTTP回复。它会遵守客户端c设置的策略（如重定向、cookie、认证）。

如果客户端的策略（如重定向）返回错误或存在HTTP协议错误时，本方法将返回该错误；**如果回应的状态码不是2xx，本方法并不会返回错误。**

如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它。如果返回值resp的主体未关闭，c下层的RoundTripper接口（一般为Transport类型）可能无法重用resp主体下层保持的TCP连接去执行之后的请求。

请求的主体，如果非nil，会在执行后被c.Transport关闭，即使出现错误。

一般应使用Get、Post或PostForm方法代替Do方法。



```go
client := &http.Client{}
	req,err := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
```



```go
func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
```

NewRequest使用指定的方法、网址和可选的主题创建并返回一个新的*Request。

如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。



### 自定义Transport

要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：

```go
tr := &http.Transport{
	TLSClientConfig:    &tls.Config{RootCAs: pool},
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

Client和Transport类型都可以安全的被多个goroutine同时使用。出于效率考虑，应该一次建立、尽量重用。



## HTTP服务端

### 默认的Server

ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。

Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

```go
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```

### 默认的Server示例

使用Go语言中的`net/http`包来编写一个简单的接收HTTP请求的Server端示例，`net/http`包是对net包的进一步封装，专门用来处理HTTP协议的数据。具体的代码如下：

```go
// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 沙河！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
```



### 自定义Server

要管理服务端的行为，可以创建一个自定义的Server：

```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```
