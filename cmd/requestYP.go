package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"reid/global"
	"reid/model"
	"strings"
)

var accessToken string
var contentType string
var headers map[string]string

type loginRequestData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// YPLoing 登陆
var userAgent string

func YPLoing() {
	w := global.GvaConfig.WebInfo
	client := resty.New()
	contentType = "application/x-www-form-urlencoded;charset=UTF-8"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	headers = map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
		"accessToken":  accessToken,
	}
	fmt.Println(headers)

	global.GvaLog.Debug("Content-Type:" + headers["Content-Type"] + "    accessToken:" + headers["accessToken"])
	post, err := client.R().
		// 将请求返回的json结果自动转换到定义的结构体类型，
		SetResult(loginRequestData{}).                                                     // 使用 post.Result().(*T).Code 打印结构体中的单个成员值
		ForceContentType("application/json").                                              //定义返回结果的类型
		SetHeaders(headers).                                                               //定义请求头
		SetBody("grant_type=password&username=" + w.Ypuser + "&password=" + w.YpPassword). //定义请求体
		Post(w.YpUrl + "/authorize/oauth2/token")
	if err != nil {
		global.GvaLog.Panic("post YP  err:", zap.Any("err", err))
	}

	//post.StatusCode()：状态码，如 200；
	//Status()：状态码和状态信息，如 200 OK；
	//Proto()：协议，如 HTTP/1.1；
	//Time()：从发送请求到收到响应的时间；
	//ReceivedAt()：接收到响应的时刻；
	//Size()：响应大小；
	//Header()：响应首部信息，以http.Header类型返回，即map[string][]string；
	//Cookies()：服务器通过Set-Cookie首部设置的 cookie 信息。
	//---------------------------------------------------
	//EnableTrace()方法启用 trace 可以记录请求的每一步的耗时和其他信息。
	//client.R().EnableTrace().Get("https://baidu.com")
	//ti := resp.Request.TraceInfo()
	//DNSLookup：DNS 查询时间，如果提供的是一个域名而非 IP，就需要向 DNS 系统查询对应 IP 才能进行后续操作；
	//ConnTime：获取一个连接的耗时，可能从连接池获取，也可能新建；
	//TCPConnTime：TCP 连接耗时，从 DNS 查询结束到 TCP 连接建立；
	//TLSHandshake：TLS 握手耗时；
	//ServerTime：服务器处理耗时，计算从连接建立到客户端收到第一个字节的时间间隔；
	//ResponseTime：响应耗时，从接收到第一个响应字节，到接收到完整响应之间的时间间隔；
	//TotalTime：整个流程的耗时；
	//IsConnReused：TCP 连接是否复用了；
	//IsConnWasIdle：连接是否是从空闲的连接池获取的；
	//ConnIdleTime：连接空闲时间；
	//RequestAttempt：请求执行流程中的请求次数，包括重试次数；
	//RemoteAddr：远程的服务地址，IP:PORT格式。

	accessToken = post.Result().(*loginRequestData).AccessToken
	if accessToken == "" {
		global.GvaLog.Panic("post YP  accessToken失败")
	}
	global.GvaLog.Info("登陆研判系统完成！AccessToken：" + accessToken)
}

// YPGetPortrait 取人像库总量
func YPGetPortrait() {
	YPLoing() // 登陆YP系统取token
	w := global.GvaConfig.WebInfo
	type T struct {
		MonitorTargetCount int `json:"monitorTargetCount"`
		RespurceInfo       struct {
			Offline int `json:"offline"`
			Online  int `json:"online"`
			Total   int `json:"total"`
		} `json:"respurceInfo"`
		TargetCount    int `json:"targetCount"`
		TargetLibBlack int `json:"targetLibBlack"`
		TargetLibRed   int `json:"targetLibRed"`
		TargetLibTotal int `json:"targetLibTotal"`
		TargetLibWhite int `json:"targetLibWhite"`
	}
	contentType = "application/json;charset=UTF-8"
	headers = map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
		"accessToken":  accessToken,
	}

	global.GvaLog.Debug("Content-Type:" + headers["Content-Type"] + "    accessToken:" + headers["accessToken"])
	client := resty.New()
	get, err := client.R().
		SetResult(T{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get(w.YpUrl + "/senseface/index/others")
	if err != nil {
		global.GvaLog.Panic("get /sensefae/index/others err:", zap.Any("err", err))
	}
	global.LogInfo.Business.PortraitCount = get.Result().(*T).TargetCount                                                      // 本月人像库总量
	global.LogInfo.Business.PortraitIncrement = global.LogInfo.Business.PortraitCount - global.GvaConfig.SysInfo.PortraitCount // 本月减上月的人像库量
	global.GvaLog.Info("获取人像库数量完成！")
}

func YPGetLog() map[string]int {

	w := global.GvaConfig.WebInfo
	d := global.LogInfo.Business.AllDay
	client := resty.New()
	//days := global.LogInfo.Business.AllDay
	get, err := client.R().
		SetResult(model.Log{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get(w.YpUrl + "/senseface/logs?page=1&number=99999&startTime=" + d[0] + "+00:00:00&endTime=" + d[len(d)-1] + "+23:59:59")
		//Get(w.YpUrl + "/senseface/logs")
	if err != nil {
		global.GvaLog.Panic("get /senseface/logs err:", zap.Any("err", err))
	}

	m := global.LogInfo.Log
	m.List = get.Result().(*model.Log).List
	global.LogInfo.Log.Total = get.Result().(*model.Log).Total

	// 将带有检索、布控、人像的日志加入到统一切片
	var alllist []string
	for _, y := range m.List {
		if strings.Contains(y.ResourceName, "检索") {
			alllist = append(alllist, y.Username+"_"+y.Realname+"_"+y.OrganName+"_检索模块")
		} else if strings.Contains(y.ResourceName, "布控") {
			alllist = append(alllist, y.Username+"_"+y.Realname+"_"+y.OrganName+"_布控模块")
		} else if strings.Contains(y.ResourceName, "人像") {
			alllist = append(alllist, y.Username+"_"+y.Realname+"_"+y.OrganName+"_人像库模块")
		} else if strings.Contains(y.ResourceName, "登录") {
			alllist = append(alllist, y.Username+"_"+y.Realname+"_"+y.OrganName+"_登录")
		}
		alllist = append(alllist, y.Username+"_"+y.Realname+"_"+y.OrganName+"_日志模块")

	}
	frequency := getCount(alllist)
	global.GvaLog.Info("获取YP系统日志完成")
	return frequency
}

// 统计切片中元素出现的次数
func getCount(s []string) map[string]int {
	frequency := make(map[string]int)
	for _, item := range s {
		fmt.Println(item)
		if frequency[item] == 0 {
			frequency[item] = 1
		} else {
			frequency[item]++
		}

	}
	return frequency
}
