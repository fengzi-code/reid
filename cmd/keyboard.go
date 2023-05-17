package cmd

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"reid/global"
)

func Keyboard(msg string) {
	fmt.Println(msg)
	_, _, err := keyboard.GetSingleKey()
	if err != nil {
		global.GvaLog.Error("退出失败")
	}

}
