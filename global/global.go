package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"reid/model"
)

var (
	GvaConfig model.YamlSetting
	GvaVp     *viper.Viper
	GvaDb     *gorm.DB
	GvaRedis  *redis.Client
	LogInfo   model.LogInfo
	GvaLog    *zap.Logger
)

// 调序单元格底色/alignment设置对齐方式、font设置字体相关、border设置边框、fill设置单元格颜色
var (
	// ExcelStyle1 alignment居中对齐方式、font设置字体宋体，14号、border设置边框为黑色，像素1
	ExcelStyle1 = `{
		"alignment":
		{"horizontal":"center","vertical":"center","wrap_text":true},
		"font":
		{"bold":false,"italic":false,"family":"宋体","size":14,"color":"#000000"},
		"border":[
		{"type":"left","color":"#000000","style":1},
		{"type":"top","color":"#000000","style":1},
		{"type":"bottom","color":"#000000","style":1},
		{"type":"right","color":"#000000","style":1}]
		}`
	// ExcelStyle2 alignment居中对齐方式、font设置字体宋体，14号、border设置边框为黑色，像素1,底色为黄色
	ExcelStyle2 = `{
		"alignment":
		{"horizontal":"center","vertical":"center","wrap_text":true},
		"font":
		{"bold":false,"italic":false,"family":"宋体","size":14,"color":"#000000"},
		"border":[
		{"type":"left","color":"#000000","style":1},
		{"type":"top","color":"#000000","style":1},
		{"type":"bottom","color":"#000000","style":1},
		{"type":"right","color":"#000000","style":1}],
		"fill":
		{"type":"pattern","color":["#FFFF37"],"pattern":1}
		}`
	ExcelStyle3 = `{
		"number_format": 10,
		"lang": "zh-cn",
		"decimal_places": 2,
		"alignment":
		{"horizontal":"center","vertical":"center"},
		"font":
		{"bold":false,"italic":false,"family":"等线","size":11,"color":"#FF0000"}
		}`
	ExcelStyle4 = `{
		"alignment":
		{"horizontal":"center","vertical":"center"},
		"font":
		{"bold":false,"italic":false,"family":"宋体","size":14,"color":"#000000"},
		"border":[
		{"type":"left","color":"#000000","style":1},
		{"type":"top","color":"#000000","style":1},
		{"type":"bottom","color":"#000000","style":1},
		{"type":"right","color":"#000000","style":1}]
		}`
)
