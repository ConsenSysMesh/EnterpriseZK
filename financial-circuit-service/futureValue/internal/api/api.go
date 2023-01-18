package api

import (
	"futurevalue/internal/api/router"
	"futurevalue/internal/package/config"
	logger "futurevalue/log"

	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func GetRouter(configPath string) *gin.Engine {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	r := router.Setup()
	return r
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()

	logger.LogPrint("Future Value ZK Circuit Running on port " + conf.Server.Port)
	logger.LogPrint("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
