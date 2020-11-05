package rdb

import (
	"fmt"
	"strconv"
)

func QueryLoginId()(int , error){

	orms,_=RegEngine()

	StrMap,err:=orms.QueryString("select * from user_login order by login_id DESC limit 1")
	if StrMap==nil||err!=nil{
		return 0, err
	}
	id,_:=strconv.Atoi(StrMap[0]["login_id"])
	fmt.Println(StrMap)
	return id,nil
}
