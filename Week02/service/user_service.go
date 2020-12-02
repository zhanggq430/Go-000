package service

import (
	"practice/Go训练营/3.第二周作业/dao"
)

type ServiceUser struct {
	Name    string
	Level   int
	Address string
	Profile string
}

func GetUserInfo(uid int) (*ServiceUser, error) {
	user := dao.Users{}
	err := user.GetById(uid)
	if err != nil {
		return nil, err
	}

	// 组装
	serviceUser := &ServiceUser{
		Name:    user.Name,
		Level:   user.Level,
		Address: user.Address,
		Profile: "其他地方获取的数据",
	}

	return serviceUser, nil
}
