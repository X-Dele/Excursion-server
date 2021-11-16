package main

import (
	"gin-vue-admin/core"
)

func main() {
	core.Viper()
	core.Zap()
	core.Database()
	db, _ := core.GetDb()
	defer db.Close()
	core.RunServer()
}
