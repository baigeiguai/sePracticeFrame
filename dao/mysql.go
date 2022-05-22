package dao

import (
	"fmt"
	"sePracticeFrame/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB =nil 

func InitMySQL(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	fmt.Println(dsn)
	DB,err:=gorm.Open("mysql",dsn)
	// DB, err := gorm.Open("mysql", "root:nishiQQ123@(81.68.250.183:3306)/gameServer?charset=utf8mb4&parseTime=True&loc=Local")
	// return 
	// MySQLClose()
	MySQLClose()
	return 
	if err != nil {
		fmt.Println("fail to connect the db!")
		return
	}
	fmt.Printf("DB: %v\n", DB)
	return 
}
func MySQLClose() {
	fmt.Printf("DB: %v\n", DB)
	DB.Close()
}
