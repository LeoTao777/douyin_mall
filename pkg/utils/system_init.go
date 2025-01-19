package utils

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	ctx = context.Background()
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	//打印配置信息
	fmt.Println("config app :", viper.Get("app"))

}

func InitMySQL() {
	//自定义模板，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //彩色
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if DB == nil {
		panic("数据库链接失败，DB为空指针")
	}
	if err != nil {
		log.Fatalf("无法连接数据库：%v", err)
	}
	err = DB.Raw("SELECT 1").Error
	if err != nil {
		log.Fatalf("无法连接到数据库服务器：%v", err)
	}
	fmt.Println("config mysql :", viper.Get("mysql"), DB)
}
