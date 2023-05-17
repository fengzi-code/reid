package cmd

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"go.uber.org/zap"
	"os"
	"reid/core"
	"reid/global"
	"strconv"
)

func ToZip() {
	var dst []string
	m := global.GvaConfig.FileInfo
	dst = append(dst, m.YpLoginSaveDir, m.AccessSaveDir, m.PcInfoSaveDir, m.ModStatSaveDir, m.YpNameSaveDir)
	fmt.Println(dst)
	Ziprachiver(dst, m.ZipSaveDir)
}

func Ziprachiver(dst []string, compressFileName string) {

	compressFileName = compressFileName + global.LogInfo.Business.Yue + "_审计.zip"
	_, err := os.Stat(compressFileName)
	if err == nil {
		err := os.Remove(compressFileName)
		if err != nil {
			global.GvaLog.Panic("删除文件失败:", zap.Any("err", err))
		} //存在即删除
	}

	z := archiver.Zip{OverwriteExisting: true, MkdirAll: false}
	err = z.Archive(dst, compressFileName)
	if err != nil {
		global.GvaLog.Panic("压缩文件失败:", zap.Any("err", err))
		fmt.Println(global.GvaVp)
	}

	global.GvaLog.Info("打包完成，文件名为：" + compressFileName)

	for _, y := range dst {
		err := os.Remove(y)
		if err != nil {
			global.GvaLog.Panic("删除文件失败:", zap.Any("err", err))
		}
	}
	v := core.LoadConfigFromYaml("portraitcount")
	v.Set("PortraitCount", global.LogInfo.Business.PortraitCount)
	err = v.WriteConfig()
	if err != nil {
		global.GvaLog.Panic("保存配置文件失败:", zap.Any("err", err))
	}
	global.GvaLog.Info("本月人像数为：" + strconv.Itoa(global.LogInfo.Business.PortraitCount) + ",更新本月的人像数量到配置文件完成！")
	Keyboard("\n\n\n按任意键退出...............")
}
