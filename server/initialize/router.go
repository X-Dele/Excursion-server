package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GvaConfig.Local.Path, http.Dir(global.GvaConfig.Local.Path)) // 为用户头像和文件提供静态地址
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)
	}
	return Router
}
