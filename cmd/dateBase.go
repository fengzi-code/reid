package cmd

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" //代码不直接使用包, 底层链接要使用!"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"reid/global"
)

// GlobalConn 创建 数据库链接句柄

func InitDB() *gorm.DB {

	// mysql: 数据库的驱动名
	// 链接数据库 --格式: 用户名:密码@协议(IP:port)/数据库名？xxx&yyy&
	//?parseTime=True&loc=Local显示北京时间
	m := global.GvaConfig.Mysql
	db, err := gorm.Open("mysql", m.Username+":"+m.Password+"@tcp("+m.Addr+")/"+m.DataBases.Database1+"?"+m.Config)
	if err != nil {
		global.GvaLog.Panic("gorm.Open err:", zap.Any("err", err))
	}
	global.GvaLog.Info("数据库连接成功！")
	// 初始连接数
	db.DB().SetMaxIdleConns(m.MaxIdleConns)
	// 最大连接数
	db.DB().SetMaxOpenConns(m.MaxOpenConns)

	// 不要复数表名
	db.SingularTable(true)

	return db

}

func InitRdis() (ClientRedis *redis.Client) {
	d := global.GvaConfig.Redis
	m := redis.Options{
		Addr:     d.Addr,
		Password: d.Password,
		DB:       d.Db,
	}

	ClientRedis = redis.NewClient(&m)

	return ClientRedis

}
