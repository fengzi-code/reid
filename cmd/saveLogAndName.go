package cmd

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"reid/global"
	"strconv"
	"strings"
)

func SaveLogAndName(frequency, ttfrequency map[string]int) {

	f, err := excelize.OpenFile(global.GvaConfig.FileInfo.YpLoginTemp)
	p, err := excelize.OpenFile(global.GvaConfig.FileInfo.AccessTemp)
	if err != nil {
		global.GvaLog.Panic("excel 打开文件失败:", zap.Any("err", err))
	}

	i := 3
	style, _ := f.NewStyle(global.ExcelStyle1)
	// 调序单元格底色/alignment设置对齐方式、font设置字体相关、border设置边框、fill设置单元格颜色
	style1, _ := f.NewStyle(global.ExcelStyle2)

	style3, _ := p.NewStyle(global.ExcelStyle1)
	// 调序单元格底色/alignment设置对齐方式、font设置字体相关、border设置边框、fill设置单元格颜色
	style4, _ := p.NewStyle(global.ExcelStyle2)
	m := global.GvaConfig.SysInfo
	//n := global.LogInfo.Log
	ii := 3
	for x, y := range frequency {
		t := FindIndex(AllUserName, strings.Split(x, "_")[0])
		if strings.Split(x, "_")[3] == "登录" {
			err = f.SetCellValue("Sheet1", "A"+strconv.Itoa(ii), ii-2)
			err = f.SetCellValue("Sheet1", "B"+strconv.Itoa(ii), m.SysName)
			err = f.SetCellValue("Sheet1", "C"+strconv.Itoa(ii), m.SysHeTongId)
			err = f.SetCellValue("Sheet1", "D"+strconv.Itoa(ii), global.LogInfo.Business.Yue)
			err = f.SetCellValue("Sheet1", "F"+strconv.Itoa(ii), strings.Split(x, "_")[1]) //realname
			err = f.SetCellValue("Sheet1", "E"+strconv.Itoa(ii), strings.Split(x, "_")[0]) //username
			err = f.SetCellValue("Sheet1", "G"+strconv.Itoa(ii), strings.Split(x, "_")[2]) // 单位名称
			err = f.SetCellValue("Sheet1", "H"+strconv.Itoa(ii), AllCityNum[t])
			err = f.SetCellValue("Sheet1", "I"+strconv.Itoa(ii), y) // 日志次数

			err = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(ii), "H"+strconv.Itoa(ii), style)
			err = f.SetCellStyle("Sheet1", "E"+strconv.Itoa(ii), "F"+strconv.Itoa(ii), style1)
			err = f.SetCellStyle("Sheet1", "I"+strconv.Itoa(ii), "I"+strconv.Itoa(ii), style1)
			if err != nil {
				global.GvaLog.Panic("excel 设置风格失败:", zap.Any("err", err))
			}
			ii += 1

		} else {
			err = p.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i-2)
			err = p.SetCellValue("Sheet1", "B"+strconv.Itoa(i), m.SysName)
			err = p.SetCellValue("Sheet1", "C"+strconv.Itoa(i), m.SysHeTongId)
			err = p.SetCellValue("Sheet1", "I"+strconv.Itoa(i), global.LogInfo.Business.Yue)
			err = p.SetCellValue("Sheet1", "E"+strconv.Itoa(i), strings.Split(x, "_")[1]) //realname
			err = p.SetCellValue("Sheet1", "D"+strconv.Itoa(i), strings.Split(x, "_")[0]) //username
			err = p.SetCellValue("Sheet1", "F"+strconv.Itoa(i), strings.Split(x, "_")[2]) // 单位名称
			err = p.SetCellValue("Sheet1", "G"+strconv.Itoa(i), AllCityNum[t])            // 单位编号
			err = p.SetCellValue("Sheet1", "H"+strconv.Itoa(i), strings.Split(x, "_")[3]) // 日志类型
			err = p.SetCellValue("Sheet1", "J"+strconv.Itoa(i), y)                        // 日志次数

			err = p.SetCellStyle("Sheet1", "A"+strconv.Itoa(i), "J"+strconv.Itoa(i), style3)
			err = p.SetCellStyle("Sheet1", "D"+strconv.Itoa(i), "E"+strconv.Itoa(i), style4)
			err = p.SetCellStyle("Sheet1", "J"+strconv.Itoa(i), "J"+strconv.Itoa(i), style4)
			i += 1
		}

	}

	if err := f.SaveAs(global.GvaConfig.FileInfo.YpLoginSaveDir); err != nil {
		global.GvaLog.Panic("excel 保存文件失败:", zap.Any("err", err))
	}

	// 图驣案件模块导出
	for x, y := range ttfrequency {
		t := FindIndex(AllUserName, strings.Split(x, "_")[0])
		fmt.Println(strings.Split(x, "_")[0], t)
		err = p.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i-2)
		err = p.SetCellValue("Sheet1", "B"+strconv.Itoa(i), m.SysName)
		err = p.SetCellValue("Sheet1", "C"+strconv.Itoa(i), m.SysHeTongId)
		err = p.SetCellValue("Sheet1", "I"+strconv.Itoa(i), global.LogInfo.Business.Yue)
		err = p.SetCellValue("Sheet1", "E"+strconv.Itoa(i), strings.Split(x, "_")[1]) //realname
		err = p.SetCellValue("Sheet1", "D"+strconv.Itoa(i), strings.Split(x, "_")[0]) //username
		err = p.SetCellValue("Sheet1", "F"+strconv.Itoa(i), strings.Split(x, "_")[2]) // 单位名称
		if t == -1 {
			err = p.SetCellValue("Sheet1", "G"+strconv.Itoa(i), "440100220000") // 单位编号
		} else {
			err = p.SetCellValue("Sheet1", "G"+strconv.Itoa(i), AllCityNum[t]) // 单位编号
		}

		err = p.SetCellValue("Sheet1", "H"+strconv.Itoa(i), strings.Split(x, "_")[3]) // 日志类型
		err = p.SetCellValue("Sheet1", "J"+strconv.Itoa(i), y)                        // 日志次数

		err = p.SetCellStyle("Sheet1", "A"+strconv.Itoa(i), "J"+strconv.Itoa(i), style3)
		err = p.SetCellStyle("Sheet1", "D"+strconv.Itoa(i), "E"+strconv.Itoa(i), style4)
		err = p.SetCellStyle("Sheet1", "J"+strconv.Itoa(i), "J"+strconv.Itoa(i), style4)
		i += 1
	}

	if err := p.SaveAs(global.GvaConfig.FileInfo.AccessSaveDir); err != nil {
		global.GvaLog.Panic("excel 保存文件失败:", zap.Any("err", err))
	}
	global.GvaLog.Info("保存用户登陆次数和日志统计完成！")
}

func FindIndex(tab []string, value string) int {
	for i, v := range tab {
		if v == value {
			return i
		}
	}
	global.GvaLog.Warn("查找用户失败:", zap.Any("err", value+",username，在研判系统找不到"))
	return -1
}
