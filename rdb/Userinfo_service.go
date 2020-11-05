package rdb

import (
	"Kit_test/types"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/common/log"
)


func InsertUser(info types.UserInfo)error{
	orms,err:=RegEngine()
	if orms==nil|| err!=nil{
		return errors.New("engine register failed or error")
	}
	_,err=orms.Insert(info)
	//_,err=orms.Exec("insert into userinfo(`id`,`name`,`user_info`,`method`) values (?,?,?,?)",)
	if err!=nil{
		return err
	}
	return nil
}

func CreateTable() {
	orms, _ := RegEngine()
	defer orms.Close()
	err := orms.CreateTables(new(types.UserInfo), new(types.UserLogin), new(types.UserResponse))
	if err != nil {
		fmt.Println("create table err=", err)
	}
}

func DeleteUser(id int)error{
	orms,_ = RegEngine()
	_,err:=orms.Where("user_id=?",id).Delete(&types.UserInfo{})
	if err!=nil{
		return err
	}
	return nil
}

func QueryAllUser()([]types.UserInfo,error){
	var ui []types.UserInfo
	var err error
	orms,err=RegEngine()
	if err!=nil{
		log.Error("dbEngine is error=", err)
		return ui,err
	}
	err=orms.Find(&ui)
	if err!=nil{
		return ui,err
	}
	return ui,nil
}

func QueryUserById(id int)(types.UserInfo,error){
	var err error
	var UI []types.UserInfo
	orms,err=RegEngine()
	if err!=nil{
		log.Error("dbEngine is error=", err)
		return types.UserInfo{},err
	}
	err=orms.ID(id).Find(&UI)
	if err!=nil{
		return types.UserInfo{},err
	}
	return UI[0],nil
}

func UpdateUser(id int,info types.UserInfo) error{
	var err error
	orms,err= RegEngine()
	if err!=nil{
		log.Error("dbEngine is error=",err)
		return err
	}
	_,err=orms.Where("user_id=?",id).Update(&info)
	if err!=nil{
		return err
	}
	return nil
}

