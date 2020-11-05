package rdb

import (
	"Kit_test/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
)

var orms *xorm.Engine

func RegEngine() (*xorm.Engine,error){
	dataSourceName:=router.User+":"+router.PASSWORD+"@tcp("+router.HOST+":"+strconv.Itoa(router.POST)+")/"+router.DATABASE+"?charset=utf8"
	//fmt.Println(dataSourceName)
	orm,err:=xorm.NewEngine("mysql", dataSourceName)
	if err!=nil{
		return nil,err
	}
	return orm,nil
}

