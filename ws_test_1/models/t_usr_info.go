package models

import "code271/ws_test_1/pkg/db"

type TUserInfo struct {
	UserID      int64  `json:"user_id,omitempty" gorm:"column:user_id;not null;primary_key;type:bigint(20)"`           // 用户id
	AccountName string `json:"account_name,omitempty" gorm:"column:account_name;not null;default '';type:varchar(20)"` // 登录名称
	Mobile      string `json:"mobile,omitempty" gorm:"column:mobile;not null;default '';type:varchar(15)"`             // 用户手机号，可用于登录
	NickName    string `json:"nick_name,omitempty" gorm:"column:nick_name;not null;default '';type:varchar(10)"`       // 用户昵称
	Password    string `json:"password,omitempty" gorm:"column:password;not null;default 'password';type:varchar(20)"` // 用户密码
	EMail       string `json:"e_mail,omitempty" gorm:"column:e_mail;not null;default '';type:varchar(30)"`             // 用户邮箱
}

func (t *TUserInfo) TableName() string {
	return `t_user_info`
}

func JudgeUserPassWord(account, password string) (ok bool, err error) {
	getOne := new(TUserInfo)
	err = db.DB.Table(getOne.TableName()).
		Where("( account_name = ? or mobile = ? or user_id = ? or nick_name = ? ) ").
		First(getOne).Error
	if err != nil {
		return
	}
	if getOne.Password == password{
		ok = true
	}
	return
}
