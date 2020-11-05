package Servers

import (
	"Kit_test/types"
	"context"
	"encoding/json"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

func DecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	switch r.Method{
	case "GET":
		id := r.URL.Query().Get("uid")
		Id, _ := strconv.Atoi(id)
		return types.UserInfo{UserId: Id}, nil
	case "POST":
		var Utype types.UserInfo
		err:=r.ParseForm()
		if err!=nil{
			log.Error(err)
		}
		for k,_:=range r.PostForm{
			//fmt.Println(k)
			_=json.Unmarshal([]byte(k),&Utype)
			//fmt.Println(Utype)
		}
		//_=rdb.InsertUser(Utype)

		return Utype,nil
	}
	return types.UserInfo{UserId: 101},nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
