package cmd

import (
	"fmt"
	"go.uber.org/zap"
	"reid/global"
	"strings"
)

var CityMap = map[string]string{
	"市局": "440100220000",
	"技术": "440100320000",
	"反恐": "440100330000",
	"公共": "440100350000",
	"荔湾": "440103170000",
	"海珠": "440105250000",
	"天河": "440106180000",
	"白云": "440111000000",
	"黄埔": "440112170000",
	"花都": "440114000000",
	"从化": "440117160000",
	"增城": "440118170000",
	"番禺": "440118170000",
}

var AllCityNum []string
var AllUserName []string

func GetUserInfo() {
	global.GvaDb = InitDB()
	err := global.GvaDb.Raw("select a.realname,a.username,a.organ_id,b.`name`,a.state FROM info_user as a,info_organization as b where a.organ_id = b.organ_id;").Scan(&global.LogInfo.UserInfo).Error
	if err != nil {
		global.GvaLog.Panic("mysql select err:", zap.Any("err", err))
	}
	for _, b := range global.LogInfo.UserInfo {
		//一个中文占用3个字节，这里取6
		citynum, ok := CityMap[strings.TrimSpace(b.Name[:6])]
		/*如果确定是真实的,则存在,否则不存在 */
		if ok {
			AllCityNum = append(AllCityNum, citynum)
		} else {
			AllCityNum = append(AllCityNum, "440100220000")
		}
		AllUserName = append(AllUserName, b.Username)
	}
	global.GvaLog.Info("从数据库获取用户完成！")
	for a, b := range global.LogInfo.UserInfo {
		fmt.Println(a, b.Username, b.Realname, b.Name, b.State, AllCityNum[a])
	}
}
