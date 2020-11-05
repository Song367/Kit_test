package hanlder

import (
	"Kit_test/model"
	"Kit_test/types"
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

func AddLoginInfo(w http.ResponseWriter,req *http.Request){
	req.ParseForm()
	var UI types.UserInfo
	var UL types.UserLogin
	for k,_:=range req.PostForm{
		_=json.Unmarshal([]byte(k),&UI)
	}

	id,err:=model.InsertLogin(UI,UL)
	if err!=nil{
		log.Error(err)
	}
	ids:=strconv.Itoa(id)
	w.Header().Set("log_id",ids)
	fmt.Println("AddLoginInfo succeed")
}

func LogoutUser(w http.ResponseWriter,req *http.Request){
	req.ParseForm()
	var UI types.UserInfo
	var UL types.UserLogin
	for k,_:=range req.PostForm{
		_=json.Unmarshal([]byte(k),&UI)
	}
	log_id,_:=strconv.Atoi(req.URL.Query().Get("log_id"))
	err:=model.LogoutUser(UI,UL,log_id)
	if err!=nil{
		log.Error("登出数据有误",err)

	}
	fmt.Println("logout succeed")
}
