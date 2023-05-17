package cmd

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"reid/global"
	"strconv"
	"time"
)

func GetRedisInfo() {

	global.GvaRedis = InitRdis()
	global.GvaLog.Info("redis 初始化完成")
	defer func(GvaRedis *redis.Client) {
		err := GvaRedis.Close()
		if err != nil {
			global.GvaLog.Panic("redis.Close err:", zap.Any("err", err))
		}
	}(global.GvaRedis)
	fmt.Println(time.Now().Year(), time.Now().AddDate(0, -1, 0).Month())
	MonthAlarmKeys, MonthCaptureKeys := monthInterval(Year, time.Now().AddDate(0, -1, 0).Month())
	fmt.Println(MonthAlarmKeys, MonthCaptureKeys)

	// 取上个月的告警量
	DaysAlarmValue, err := global.GvaRedis.MGet(MonthAlarmKeys...).Result()
	if err != nil {
		global.GvaLog.Panic("redis.get err:", zap.Any("err", err))
	}
	fmt.Println("每日告警量：", DaysAlarmValue)
	for _, y := range DaysAlarmValue {
		if y != nil {
			DaysAlarm, _ := strconv.Atoi(y.(string))
			global.LogInfo.Business.LastMonthAlarm += DaysAlarm
		}
	}
	fmt.Println("告警量：", global.LogInfo.Business.LastMonthAlarm)

	// 取上个月的抓拍量
	DaysAlarmValue, _ = global.GvaRedis.MGet(MonthCaptureKeys...).Result()
	fmt.Println("每日抓拍量：", DaysAlarmValue)
	for _, y := range DaysAlarmValue {
		if y != nil {
			DaysAlarm, _ := strconv.Atoi(y.(string))
			global.LogInfo.Business.LastCapture += DaysAlarm
		}
	}
	fmt.Println("抓拍量：", global.LogInfo.Business.LastCapture)
	global.GvaLog.Info("读取抓拍量和告警量完成！")
}
