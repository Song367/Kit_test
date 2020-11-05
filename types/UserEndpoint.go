package types

import "time"

type UserInfo struct{
	UserId 			int 				`json:"user_id" xorm:"user_id" db:"user_id"`
	UserName 		string 				`json:"username" xorm:"username" db:"username"`
	Password 		string 				`json:"password" xorm:"password" db:"paaword"`
	Sex 			string 				`json:"sex" xorm:"sex" db:"sex"`
	CreateTime 		time.Time 			`json:"create_time" xorm:"create_time" db:"create_time"`
}

type UserLogin struct{
	LoginUserId    	int				`json:"login_user_id" db:"login_user_id" xorm:"login_user_id"`
	LoginTime 		time.Time		`json:"login_time" db:"login_time" xorm:"login_time"`
	LogoutTime 		time.Time 		`json:"logout_time" db:"logout_time" xorm:"logout_time"`
	LoginStatus 	uint     		`json:"login_status" db:"login_status" xorm:"login_status"`  // 1:在线 2:离线 3:勿打扰
	LoginId 		int				`json:"login_id" db:"login_id" xorm:"login_id"`

}

type UserResponse struct {
	Result 			interface{} 	`json:"result"`
}
