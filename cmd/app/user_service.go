package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	// echoOpenAPI "github.com/alexferl/echo-openapi"
	httpHandler "github.com/fahmyabida/eDot/pkg/http/handler"
	customMiddleware "github.com/fahmyabida/eDot/pkg/http/middleware"
)

var userCommand = &cobra.Command{
	Use:   "user",
	Short: "Start User server",
	Run:   RunUserServer,
}

func init() {
	rootCmd.AddCommand(userCommand)
}

func RunUserServer(cmd *cobra.Command, args []string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(customMiddleware.ErrorMiddleware())

	// e.Use(echoOpenAPI.OpenAPI("./docs/openapi.yaml"))

	healthcheckGroup := e.Group("/healthcheck")
	httpHandler.InitHealthcheckHandler(healthcheckGroup)

	v1 := e.Group("/api/v1")
	v1NoAuth := e.Group("/api/v1")
	v1.Use(customMiddleware.AuthMiddleware(jwtConfig))
	httpHandler.InitUserHandler(v1, v1NoAuth, UserUsecase)

	httpHandler.InitProductHandler(v1, ProductUsecase)
	httpHandler.InitOrderHandler(v1, OrderUsecase)
	httpHandler.InitWarehouseHandler(v1, WarehouseUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
