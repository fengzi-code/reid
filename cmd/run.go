package cmd

import (
	"fmt"
	"reid/global"
	"time"
)

var Year int

func Run() {
	global.GvaLog.Info("------------------开始运行------------------")
	m := fmt.Sprintf("%d", time.Now().Month())
	if m == "1" {
		Year = time.Now().Year() - 1
	} else {
		Year = time.Now().Year()
	}
	fmt.Println(Year)
	global.LogInfo.Business.Yue = fmt.Sprintf("%d%d", Year, time.Now().AddDate(0, -1, 0).Month())
	fmt.Println(global.LogInfo.Business.Yue)
	GetUserInfo()             //从数据库取用户数据
	GetRedisInfo()            //从redis取抓拍和告警
	YPGetPortrait()           //取人像库上月数量
	SaveSysname()             //导出系统未禁用用户到excel表
	frequency := YPGetLog()   // 从YP系统取日志列表
	ttfrequency := GetTTLog() // 取TT系统的日志列表
	SaveLogAndName(frequency, ttfrequency)
	SavePcInfo(global.GvaConfig.FileInfo.PcInfoTemp, global.GvaConfig.FileInfo.PcInfoSaveDir)
	SaveModStat(global.GvaConfig.FileInfo.ModStatTemp, global.GvaConfig.FileInfo.ModStatSaveDir)
	ToZip()
}
