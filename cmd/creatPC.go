package cmd

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"math/rand"
	"reid/global"
	"strconv"
	"time"
)

func SavePcInfo(PcInfoTemp, PcInfoSaveDir string) {

	f, err := excelize.OpenFile(PcInfoTemp)
	if err != nil {
		global.GvaLog.Panic("excel 打开文件失败:", zap.Any("err", err))
	}

	// 调序单元格底色/alignment设置对齐方式、font设置字体相关、border设置边框、fill设置单元格颜色
	style1, _ := f.NewStyle(global.ExcelStyle3)

	for i := 2; i < 21; i++ {
		if i < 19 {
			fmt.Println(i)
			PcCpuTop, _ := f.GetCellValue("Sheet1", "I"+strconv.Itoa(i))
			PcCpuAve, _ := f.GetCellValue("Sheet1", "J"+strconv.Itoa(i))
			RamTop, _ := f.GetCellValue("Sheet1", "K"+strconv.Itoa(i))
			RamAve, _ := f.GetCellValue("Sheet1", "L"+strconv.Itoa(i))
			DiskUse, _ := f.GetCellValue("Sheet1", "M"+strconv.Itoa(i))
			PcCpuTopFloat, _ := strconv.ParseFloat(PcCpuTop, 64)
			PcCpuAveFloat, _ := strconv.ParseFloat(PcCpuAve, 64)
			RamTopFloat, _ := strconv.ParseFloat(RamTop, 64)
			RamAveTopFloat, _ := strconv.ParseFloat(RamAve, 64)
			DiskUseFloat, _ := strconv.ParseFloat(DiskUse, 64)
			// 置随机数种子，取随机浮点数
			rand.Seed(time.Now().Unix())
			// 保留4位浮点数
			PcCpuTopFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", PcCpuTopFloat-rand.Float64()/35), 64)
			PcCpuAveFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", PcCpuAveFloat-rand.Float64()/35), 64)
			RamTopFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", RamTopFloat-rand.Float64()/35), 64)
			RamAveTopFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", RamAveTopFloat-rand.Float64()/35), 64)
			DiskUseFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", DiskUseFloat-rand.Float64()/35), 64)

			fmt.Println(PcCpuTopFloat, PcCpuAveFloat, RamTopFloat, RamAveTopFloat, DiskUseFloat)
			err = f.SetCellStyle("Sheet1", "I"+strconv.Itoa(i), "M"+strconv.Itoa(i), style1)
			err := f.SetCellValue("Sheet1", "I"+strconv.Itoa(i), PcCpuTopFloat)
			err = f.SetCellValue("Sheet1", "J"+strconv.Itoa(i), PcCpuAveFloat)
			err = f.SetCellValue("Sheet1", "K"+strconv.Itoa(i), RamTopFloat)
			err = f.SetCellValue("Sheet1", "L"+strconv.Itoa(i), RamAveTopFloat)
			err = f.SetCellValue("Sheet1", "M"+strconv.Itoa(i), DiskUseFloat)
			if err != nil {
				global.GvaLog.Panic("excel 设置失败:", zap.Any("err", err))
			}

		} else {

			DiskUse, _ := f.GetCellValue("Sheet1", "M"+strconv.Itoa(i))
			DiskUseFloat, _ := strconv.ParseFloat(DiskUse, 64)
			DiskUseFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", DiskUseFloat-rand.Float64()/35), 64)
			err = f.SetCellStyle("Sheet1", "M"+strconv.Itoa(i), "M"+strconv.Itoa(i), style1)
			err = f.SetCellValue("Sheet1", "M"+strconv.Itoa(i), DiskUseFloat)
		}
		if err := f.SaveAs(PcInfoSaveDir); err != nil {
			global.GvaLog.Panic("excel 保存文件失败:", zap.Any("err", err))
		}
	}
	global.GvaLog.Info("保存系统硬件资源使用情况完成！")
}
