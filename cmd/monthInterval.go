package cmd

import (
	"reid/global"
	"strconv"
	"strings"
	"time"
)

func monthInterval(y int, m time.Month) (MonthAlarmKeys, MonthCaptureKeys []string) {
	r := global.GvaConfig.Redis
	firstDay := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	lastDay := time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
	lastDayInt := strings.Split(lastDay.Format("2006-01-02"), "-")
	EndNum, _ := strconv.Atoi(lastDayInt[len(lastDayInt)-1]) //取最后一天
	//EndMot, _ := strconv.Atoi(lastDayInt[len(lastDayInt)-2])    //取月份

	var MonthAlarmKey string
	var MonthCaptureKey string
	var OneDay string
	var Allday []string
	//var DayValues []interface{}
	for i := 1; i < EndNum+1; i++ {
		if i < 10 {
			OneDay = firstDay.Format("2006-01") + "-0" + strconv.Itoa(i)
		} else {
			OneDay = firstDay.Format("2006-01") + "-" + strconv.Itoa(i)
		}
		MonthAlarmKey = r.MonthAlarmKey + OneDay
		MonthCaptureKey = r.MonthCaptureKey + OneDay
		MonthAlarmKeys = append(MonthAlarmKeys, MonthAlarmKey)
		MonthCaptureKeys = append(MonthCaptureKeys, MonthCaptureKey)
		// 返回 DayValues []interface{} 类型给 ClientRedis.MSet(DayValue...)批量设置健值
		//DayValues = append(DayValues,MonthCaptureKey)
		//DayValues = append(DayValues,2436643)
		// 把所有日期加入切片
		Allday = append(Allday, OneDay)
	}
	global.LogInfo.Business.AllDay = Allday
	return MonthAlarmKeys, MonthCaptureKeys
}
