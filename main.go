package main

import (
	"Kit_test/Servers"
	"Kit_test/hanlder"
	httppoint "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	user:=Servers.UserServer{}
	ept:= Servers.GetUserInfo(&user)
	server:=httppoint.NewServer(ept,Servers.DecodeRequest,Servers.EncodeResponse)

	route:=mux.NewRouter()
	{
		route.Methods("POST").Path("/index").Handler(server)

		route.Methods("POST").Path("/insert_user").HandlerFunc(hanlder.RegisterUser)
		route.Methods("GET").Path("/remove_user").HandlerFunc(hanlder.RemoveUser)
		route.Methods("POST").Path("/QueryAll").HandlerFunc(hanlder.QueryAll)
		route.Methods("GET").Path("/queryuserbyid").HandlerFunc(hanlder.QueryUserById)
		route.Methods("POST").Path("/update").HandlerFunc(hanlder.UpdateUser)
	}
	{
		route.Methods("POST").Path("/login").HandlerFunc(hanlder.AddLoginInfo)
		route.Methods("POST").Path("/logout").HandlerFunc(hanlder.LogoutUser)
	}
	_=http.ListenAndServe(":8080",route)
}
