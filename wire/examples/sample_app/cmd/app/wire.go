//go:build wireinject
// +build wireinject

package main

import (
	"http_app/pkg/mysql"
	"http_app/pkg/redis"

	"github.com/google/wire"
)

func CreateApplication() Application {
	wire.Build(NewApplication, redis.NewRedis, mysql.NewMySQL)
	return Application{}
}
