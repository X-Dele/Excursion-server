package core

import (
	"database/sql"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gorm.io/gorm"
)

func Database() {
	global.GvaDb = initialize.Gorm()
	initTables(global.GvaDb)
}

func GetDb() (*sql.DB, error) {
	return global.GvaDb.DB()
}

func initTables(db *gorm.DB) {
	initialize.MysqlTables(db)
}
