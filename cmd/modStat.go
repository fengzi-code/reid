package cmd

import (
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"reid/core"
	"reid/global"
	"strconv"
)

func SaveModStat(ModStatTemp, ModStatSaveDir string) {

	f, err := excelize.OpenFile(ModStatTemp)
	if err != nil {
		global.GvaLog.Panic("excel 打开文件失败:", zap.Any("err", err))
	}

	style3, _ := f.NewStyle(global.ExcelStyle4)

	for i := 3; i < 10; i++ {
		err = f.SetCellStyle("Sheet1", "D"+strconv.Itoa(i), "D"+strconv.Itoa(i), style3)
		err := f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), global.LogInfo.Business.Yue)
		if err != nil {
			global.GvaLog.Panic("excel 设置文件失败:", zap.Any("err", err))
		}
	}
	m := global.LogInfo
	err = f.SetCellValue("Sheet1", "G3", m.Log.Total) //# 检索模块	检索量 --日志总数
	v := core.LoadConfigFromYaml("portraitcount")     // 从配置文件里取上个月的人像数
	PortraitIncrement := m.Business.PortraitCount - v.Get("portraitCount").(int)
	err = f.SetCellValue("Sheet1", "G4", PortraitIncrement)         //# 人像库模块	新建量--新增人像数
	err = f.SetCellValue("Sheet1", "G5", m.Business.PortraitCount)  //# 人像库模块	保存量--现有人像数
	err = f.SetCellValue("Sheet1", "G6", m.Business.LastCapture)    //# 布控模块	比对量--抓拍量
	err = f.SetCellValue("Sheet1", "G7", m.Business.LastMonthAlarm) //# 布控模块	告警量--
	err = f.SetCellValue("Sheet1", "G8", m.Business.TTAnJianCount)  //# 案件模块	新增量--图腾
	err = f.SetCellValue("Sheet1", "G9", m.TTLog.Total)             //# 日志模块	新增量--图腾

	if err := f.SaveAs(ModStatSaveDir); err != nil {
		global.GvaLog.Panic("excel 保存文件失败:", zap.Any("err", err))
	}
	global.GvaLog.Info("保存系统业务量完成！")
}
