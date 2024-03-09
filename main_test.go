package main

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"github.com/MrMohebi/golang-gin-boilerplate.git/router"
	"github.com/gin-gonic/gin"
)

// nodemon --exec go test main_test.go --signal SIGTERM

func main() {
	configs.Setup()

	server := gin.Default()

	router.Routs(server)

	err := server.Run(configs.EnvServeHost())
	common.IsErr(err, true, "Err in starting server")
}
