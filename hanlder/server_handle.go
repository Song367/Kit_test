package hanlder

import (
	"Kit_test/model"
	"Kit_test/types"
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
	"time"
)


func RegisterUser(writer http.ResponseWriter,request *http.Request) {
	_=request.ParseForm()
	var Info types.UserInfo
	for k,_:=range request.PostForm{
		_=json.Unmarshal([]byte(k),&Info)
	}
	Info.CreateTime = time.Now()
	//ctx,_:=context.WithCancel(context.Background())
	err:=model.RegisterUser(Info)
	if err!=nil{
		log.Error("insert user err=",err)
	}
	fmt.Println("Successful for insert")
}

func RemoveUser(writer http.ResponseWriter,request *http.Request){
	r:=request.URL.Query().Get("uid")
	id,_:=strconv.Atoi(r)
	err:=model.RemoveUser(id)
	if err!=nil{
		log.Error("Remove user",err)
	}
	fmt.Println("Deletion succeed")
}

func QueryAll(writer http.ResponseWriter , req *http.Request){

	UI,err:=model.QueryAllUser()
	if err!=nil{
		log.Error(err)
	}
	fmt.Println(UI)
	var Ubyte []byte
	Ubyte,_=json.Marshal(UI)

	writer.Header().Set("Content-type","form-data")
	writer.Write(Ubyte)
}

func QueryUserById(writer http.ResponseWriter , req *http.Request){
	ids:=req.URL.Query().Get("uid")
	id,_:=strconv.Atoi(ids)
	UI,err:=model.QueryUserById(id)
	if err!=nil{
		log.Error("根据用户id查询",err)
	}
	Mar,_:=json.Marshal(&UI)
	writer.Header().Set("Content-type","form/data")
	_,_=writer.Write(Mar)
}

func UpdateUser(writer http.ResponseWriter , req *http.Request){
	_=req.ParseForm()
	var ui types.UserInfo
	for k,_:=range req.PostForm{
		_=json.Unmarshal([]byte(k),&ui)
	}
	err:=model.UpdateUser(ui.UserId,ui)
	if err!=nil{
		log.Error("UpdateUser err",err)
	}
	fmt.Println("UpdateUser succeed")
}

