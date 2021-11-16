package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.GET("push_url", v1.GetPushUrl)
	}
	return BaseRouter
}
