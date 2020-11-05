package Servers

type OpUserServer interface {
	GetInfo(id int) string
}

type UserServer struct{}

func (Self *UserServer)GetInfo(id int)string{
	if id==101{
		return "LeBron"
	}else{
		return "DWade"
	}
}
