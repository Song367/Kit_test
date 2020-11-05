package rdb

import (
	"Kit_test/types"
	"github.com/prometheus/common/log"
)

func AddLoginInfo(login types.UserLogin)error{
	var err error
	orms,err=RegEngine()
	if err!=nil{
		log.Error("dbEngine err",err)
		return err
	}
	_,err=orms.Insert(&login)
	if err!=nil{
		return err
	}

	return nil
}

func LogoutUser(login types.UserLogin,Login_id int)error{
	var err error

	orms,err=RegEngine()
	if err!=nil{
		log.Error("dbEngine err",err)
		return err
	}
	_,err=orms.Where("login_id=?",Login_id).Update(&login)
	if err!=nil{
		return err
	}
	return nil
}


