package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
	)
	if err != nil {
		global.GvaLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GvaLog.Info("register table success")
}

func Gorm() *gorm.DB {
	return gormMysql()
}

// GormMysql 初始化Mysql数据库
func gormMysql() *gorm.DB {
	m := global.GvaConfig.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.GvaLog.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	if global.GvaConfig.Mysql.LogZap {
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	if mod {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}
