package main

import (
	"cashier-api/core/config"
	"cashier-api/core/server"
	"flag"
	"fmt"
	"io"

	"cashier-api/docs"
	"cashier-api/middleware"
	"cashier-api/router"

	"cashier-api/core/logger"
	"cashier-api/core/translator"

	"github.com/labstack/gommon/log"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// -----------
	// Config
	// -----------
	configPath := flag.String("config", "configs", "set config path")
	environment := flag.String("environment", "local", "set environment")
	flag.Parse()
	err := config.Read(*configPath, *environment)
	if err != nil {
		panic(err)
	}

	err = config.ReadReturnResult(*configPath, "return_results")
	if err != nil {
		panic(err)
	}

	translator.InitTranslator()

	options := &router.Options{
		Environment: config.ENV,
		Results:     config.RR,
	}

	// programatically set swagger info
	docs.SwaggerInfo.Title = config.ENV.Swagger.Title
	docs.SwaggerInfo.Description = config.ENV.Swagger.Description
	docs.SwaggerInfo.Version = config.ENV.Swagger.Version
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", config.ENV.Swagger.Host, config.ENV.Swagger.BaseURL)
	//=======================================================

	// -----------
	// Log
	// -----------
	var logLevel log.Lvl
	var logHeader string
	var logOutput io.Writer

	if config.ENV.App.Release {
		logLevel = log.INFO
		options.LogMiddleware = middleware.Logger()
	} else {
		logLevel = log.DEBUG
		logHeader = "\033[1;34m-->\033[0m ${time_rfc3339} ${level}"
		options.LogMiddleware = middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "\033[1;34m-->\033[0m method=${method} \033[1;32muri=${uri}\033[0m user_agent=${user_agent} " +
				"status=${status} error=${error} latency_human=${latency_human}, \n\033[1;93mparameters=${parameters}\nuser_id=${user_id}\033[0m\n",
		})
	}
	// -----------
	// Router
	// -----------
	e := router.NewWithOptions(options)
	if logOutput != nil {
		e.Logger.SetOutput(logOutput)
	}
	e.Logger.SetPrefix("Cashier")
	e.Logger.SetLevel(logLevel)
	if logHeader != "" {
		e.Logger.SetHeader(logHeader)
	}
	logger.Logger = e.Logger

	server.New(e, config.ENV.App.Port).Start()
}
