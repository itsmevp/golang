package main

import (
	"fmt"
	"http_app/pkg/mysql"
	"http_app/pkg/redis"
)

type Application struct {
	mysql mysql.MySQL
	redis redis.Redis
}

func NewApplication(m mysql.MySQL, r redis.Redis) Application {
	return Application{mysql: m, redis: r}
}

func main() {
	app := CreateApplication()
	fmt.Println(app)
}
