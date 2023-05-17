package cmd

import (
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"reid/global"
	"strconv"
)

func SaveSysname() {
	f, err := excelize.OpenFile(global.GvaConfig.FileInfo.YpNameTemp)
	if err != nil {
		global.GvaLog.Panic("redis.Close err:", zap.Any("err", err))

	}

	i := 3
	style, _ := f.NewStyle(global.ExcelStyle1)
	// 调序单元格底色/alignment设置对齐方式、font设置字体相关、border设置边框、fill设置单元格颜色
	style1, _ := f.NewStyle(global.ExcelStyle2)
	for x, y := range global.LogInfo.UserInfo {
		m := global.GvaConfig.SysInfo
		if y.State == 1 {
			err := f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i-2)
			if err != nil {
				global.GvaLog.Panic("excel 写入失败:", zap.Any("err", err))
			}
			err = f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), m.SysName)
			err = f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), m.SysHeTongId)
			err = f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), global.LogInfo.Business.Yue)
			err = f.SetCellValue("Sheet1", "F"+strconv.Itoa(i), y.Realname)
			err = f.SetCellValue("Sheet1", "E"+strconv.Itoa(i), y.Username)
			err = f.SetCellValue("Sheet1", "G"+strconv.Itoa(i), y.Name)
			err = f.SetCellValue("Sheet1", "H"+strconv.Itoa(i), AllCityNum[x])

			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i), "H"+strconv.Itoa(i), style)
			if err != nil {
				global.GvaLog.Panic("excel 设置风格失败:", zap.Any("err", err))
			}
			err = f.SetCellStyle("Sheet1", "E"+strconv.Itoa(i), "F"+strconv.Itoa(i), style1)
			if err != nil {
				global.GvaLog.Panic("excel 设置风格失败:", zap.Any("err", err))
			}
			i += 1

		}
	}
	s := global.GvaConfig.FileInfo.YpNameSaveDir
	if err := f.SaveAs(s); err != nil {
		global.GvaLog.Panic("excel 保存文件失败:", zap.Any("err", err))
	}
	global.GvaLog.Info("保存研判系统未禁用用户完成！")
}
