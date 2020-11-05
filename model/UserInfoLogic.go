package model

import (
	"Kit_test/common"
	"Kit_test/rdb"
	_ "Kit_test/rdb"
	"Kit_test/types"
	_ "fmt"
)

func RegisterUser(info types.UserInfo) error{
 	res:=common.Encryption(info.Password)
 	info.Password = res
 	err:= rdb.InsertUser(info)
	if err!=nil{
		return err
	}
	return nil
}

func RemoveUser(id int)error{
	err:=rdb.DeleteUser(id)
	if err!=nil{
		return err
	}
	return nil
}

func QueryAllUser()([]types.UserInfo,error){
	var UI []types.UserInfo
	var err error
	UI,err=rdb.QueryAllUser()
	if err!=nil{
		return UI,err
	}
	return UI,nil
}

func QueryUserById(id int)(interface{},error){
	var UI types.UserInfo
	var err error
	UI,err=rdb.QueryUserById(id)
	if err != nil {
		return nil,err
	}
	return UI,nil
}

func UpdateUser(id int,info types.UserInfo)error{
	if info.Password==""{
		err:=rdb.UpdateUser(id,info)
		if err!=nil{
			return err
		}
		return nil
	}
	pwd := common.Encryption(info.Password)
	info.Password = pwd
	err:=rdb.UpdateUser(id,info)
	if err!=nil{
		return err
	}
	return nil
}

