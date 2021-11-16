package core

import (
	"gin-vue-admin/initialize"
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunServer() {
	Router := initialize.Routers()
	s := initServer(":8080", Router)
	s.ListenAndServe()
}
