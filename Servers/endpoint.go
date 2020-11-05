package Servers

import (
	"Kit_test/types"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

func GetUserInfo(server OpUserServer)endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		r :=request.(types.UserInfo)
		fmt.Println(r.UserId)
		res:=server.GetInfo(r.UserId)
		return types.UserResponse{Result: res},nil
	}
}
