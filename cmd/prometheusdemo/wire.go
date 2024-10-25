//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"prometheusdemo/handler"
	"prometheusdemo/ioc"
)

func InitWeb() *gin.Engine {
	panic(wire.Build(ioc.InitConfig, ioc.InitDB, ioc.InitWebServer, ioc.InitMiddlewares, handler.NewUserHandler))
}
