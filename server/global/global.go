package global

import (
	"gin-vue-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GvaLog *zap.Logger
	GvaDb  *gorm.DB
	GvaConfig config.Server
)
