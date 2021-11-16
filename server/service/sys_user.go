package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
)



//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GvaDb.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}
