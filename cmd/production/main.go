package main

import (
	"github.com/labstack/echo/v4"
	"home-test-tiki/handler"
	"home-test-tiki/log"
	"home-test-tiki/router"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "backend_testing_dev")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	// loads values from .env into the system
	//if err := godotenv.Load("config_file/.env"); err != nil {
	//	log.Print("No .env file found")
	//}
}

func main() {
	e := echo.New()

	handlerSplitString :=  handler.StringHandler{

	}

	api := router.API{
		Echo:                 e,
		StringHandler: handlerSplitString,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":80"))
}
