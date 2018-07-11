package main

import (
  "os"
  "fmt"

  "wechat-miniprogram/application"
  "wechat-miniprogram/utils/server"
  "wechat-miniprogram/utils/database"
)

func main() {
  app := application.App{}

  serverConfig, err := server.ReadConfig("./serverConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

  dbConfig, err := database.ReadConfig("./dbConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

  app.InitializeApp(dbConfig, serverConfig)
  app.Run()
  app.Logger.Log(
    application.LOG_LAYER_TAG, application.LAYER_APPLICATION,
    application.LOG_MESSAGE_TAG, application.MESSAGE_HALTING,
    application.LOG_ERROR_TAG, <-app.Errs)
}
