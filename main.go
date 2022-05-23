package main

import (
	"fmt"
	"os"
	"sePracticeFrame/dao"
	"sePracticeFrame/models"
	"sePracticeFrame/routers"
	"sePracticeFrame/settings"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage ./main conf/config.ini")
		return
	}
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed,err:%v\n", err)
	}
	err := dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("fail to init MySQL...")
		return
	}
	dao.DB.AutoMigrate(&models.User{})
	r := routers.SetupRouters()
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Println("server fails to start while running the port!")
	}
}
