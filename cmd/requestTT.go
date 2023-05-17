package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"reid/global"
	"reid/model"
	"strings"
)

func LoingTT() (accessToken string) {

	type T struct {
		AccessToken string `json:"accessToken"`
		Expires     int    `json:"expires"`
		First       int    `json:"first"`
		OrgId       int    `json:"orgId"`
		OrgName     string `json:"orgName"`
		Other       struct {
		} `json:"other"`
		Permission struct {
		} `json:"permission"`
		PermissionIds []int  `json:"permissionIds"`
		Realname      string `json:"realname"`
		Roles         []int  `json:"roles"`
		UserId        int    `json:"userId"`
		Username      string `json:"username"`
	}
	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	w := global.GvaConfig.WebInfo
	client := resty.New()
	post, err := client.R().
		// 将请求返回的json结果自动转换到定义的结构体类型，
		SetResult(T{}).                                            // 使用 post.Result().(*T).Code 打印结构体中的单个成员值
		ForceContentType("application/json").                      //定义返回结果的类型
		SetHeaders(headers).                                       //定义请求头
		SetBody(user{Username: w.TTuser, Password: w.TTPassword}). //定义请求体
		//SetBody(`{"username":"sfadmin2","password":"sense2018"}`).
		Post(w.TTUrl + "/uums/login")
	if err != nil {
		global.GvaLog.Panic("post /uums/login 失败:", zap.Any("err", err))
	}
	//global.GvaLog.Info("info:",zap.Any("info",w.YpUrl + "/uums/login"),zap.Any("headers",user{Username: w.TTuser, Password: w.TTPassword}))
	accessToken = post.Result().(*T).AccessToken
	if accessToken == "" {
		global.GvaLog.Panic("post TT  accessToken失败")
	}
	global.GvaLog.Info("登陆图腾系统完成！accessToken: " + accessToken)
	fmt.Println(accessToken)
	return accessToken

}

func GetTTLog() map[string]int {
	accessToken = LoingTT()
	headers = map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
		"accessToken":  accessToken,
	}
	//days := global.LogInfo.Business.AllDay
	w := global.GvaConfig.WebInfo
	d := global.LogInfo.Business.AllDay
	client := resty.New()
	get, err := client.R().
		// 将请求返回的json结果自动转换到定义的结构体类型，
		SetResult(model.TTLog{}).             // 使用 post.Result().(*T).Code 打印结构体中的单个成员值
		ForceContentType("application/json"). //定义返回结果的类型
		SetHeaders(headers).                  //定义请求头
		Get(w.TTUrl + "/totem/totem-logs/logs?number=99999&type=1&page=1&sort=desc&order=time&start=" + d[0] + "&end=" + d[len(d)-1])
	if err != nil {
		global.GvaLog.Panic("get /totem/totem-logs 失败:", zap.Any("err", err))
	}
	m := global.LogInfo.TTLog
	m.List = get.Result().(*model.TTLog).List
	global.LogInfo.TTLog.Total = get.Result().(*model.TTLog).Total
	// 将带有检索、布控、人像的日志加入到统一切片
	var alllist []string

	for _, y := range m.List {
		if strings.Contains(y.ResourceName, "案件") {
			alllist = append(alllist, y.Username+"_"+y.Realname+"_"+y.OrganName+"_案件模块")
			global.LogInfo.Business.TTAnJianCount += 1
		}

	}
	if global.LogInfo.Business.TTAnJianCount == 0 {
		global.GvaLog.Error("get log 0 :", zap.Any("get", get), zap.Any("headers:", headers))
	}
	frequency := getCount(alllist)
	global.GvaLog.Info("获取图腾系统日志完成！")
	return frequency
}
