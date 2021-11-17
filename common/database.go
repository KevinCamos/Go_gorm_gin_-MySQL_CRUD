package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"starbars/Config"
/* 	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os" */
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB
var err error

func Init() *gorm.DB {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	fmt.Println("Status:", "entra")
	return Config.DB
/* 	defer Config.DB.Close()
 */



 /* Thinkster -> */

	// db, err := gorm.Open("sqlite3", "./../gorm.db")
	// if err != nil {
	// 	fmt.Println("db err: (Init) ", err)
	// }
	// db.DB().SetMaxIdleConns(10)
	// //db.LogMode(true)
	// DB = db
	// return DB

	 /* <-Thinkster */

}