package core

import (
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

func Zap() () {
	global.GvaLog = initialize.InitZap()
}
