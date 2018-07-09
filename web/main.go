package main

import (
	_ "OpsOauth2/web/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)


func init() {
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbName := os.Getenv("DB_NAME")
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", uri)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal("Create table error:", err)
	}

}

func main() {
	beego.SetStaticPath("/static","static")
	beego.Run()
}
