package main

import (
	"e-marketing/pkg/logger"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	initConfig()

	app := InitApp()

	// 启动定时任务，不要立即停止
	app.cron.Start()

	// 启动Web服务器
	logger.Logger.Error("start server failed ", zap.Error(app.server.Run(":8081")))

	// 只有当服务器关闭后才停止cron
	app.cron.Stop()
}

func initConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
