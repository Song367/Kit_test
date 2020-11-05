package model

import (
	"Kit_test/common"
	"Kit_test/rdb"
	"Kit_test/types"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"time"
)

func InsertLogin(ui types.UserInfo,login types.UserLogin)(int,error){
	// 判断用户是否纯在



	// 判断密码
	pwd:=common.Encryption(ui.Password)
	user,err:=QueryUserById(ui.UserId)
	if err!=nil{
		log.Error("获取密码有误")
		return 0, err
	}
	UI:=user.(types.UserInfo)
	if UI.Password!=pwd{
		log.Error("输入密码有误")
		return -1,errors.New("输入密码有误")
	}
	login.LoginUserId = ui.UserId
	login.LoginTime = time.Now()
	login.LoginStatus = 1
	id,err:=rdb.QueryLoginId()
	if err!=nil{
		log.Error("查询长度出错")
		return 0,err
	}
	login.LoginId = id+1
	fmt.Println(id)
	err=rdb.AddLoginInfo(login)
	if err!=nil{
		return 0,err
	}

	return id+1,nil
}

func LogoutUser(ui types.UserInfo,login types.UserLogin,Login_id int)error{
	login.LoginUserId = ui.UserId
	login.LogoutTime = time.Now()
	login.LoginStatus = 2
	login.LoginId = Login_id
	err:=rdb.LogoutUser(login,Login_id)
	if err!=nil{
		return err
	}
	return nil
}
