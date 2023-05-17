package main

import (
	"reid/cmd"
	"reid/core"
	"reid/global"
)

//# 检索模块	检索量 --日志总数
//# 人像库模块	新建量--新增人像数
//# 人像库模块	保存量--现有人像数
//# 布控模块	比对量--抓拍量
//# 布控模块	告警量--
//# 案件模块	新增量--图腾
//# 日志模块	新增量--图腾
// select * from senseface.Operate_Log where Operate_Time like '2021-11-%' ;
//select t.realname,t.username,t.state,t.organ_id,d.name from info_user as t,info_organization as d where t.organ_id = d.organ_id;
func main() {
	global.GvaVp = core.LoadConfigFromYaml("config")
	global.GvaLog = core.Zap() // 初始化zap日志库
	cmd.Run()
}
