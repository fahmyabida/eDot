package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	// echoOpenAPI "github.com/alexferl/echo-openapi"
	httpHandler "github.com/fahmyabida/eDot/pkg/http/handler"
	customMiddleware "github.com/fahmyabida/eDot/pkg/http/middleware"
)

var warehouseCommand = &cobra.Command{
	Use:   "warehouse",
	Short: "Start Warehouse server",
	Run:   RunWarehouseServer,
}

func init() {
	rootCmd.AddCommand(warehouseCommand)
}

func RunWarehouseServer(cmd *cobra.Command, args []string) {
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
	v1.Use(customMiddleware.AuthMiddleware(jwtConfig))
	httpHandler.InitWarehouseHandler(v1, WarehouseUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
