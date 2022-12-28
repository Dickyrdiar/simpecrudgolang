package main

import (
	"github.com/dickyrdiar/go-gin-api-tasks/pkg/books"
	"github.com/dickyrdiar/go-gin-api-tasks/pkg/common/db"
	"github.com/dickyrdiar/go-gin-api-tasks/pkg/users"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbURL := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbURL)

	books.RegisRoutes(r, h)
	users.UserRoutes(r, h)

	r.Run(port)
}
