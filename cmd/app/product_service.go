package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	// echoOpenAPI "github.com/alexferl/echo-openapi"
	httpHandler "github.com/fahmyabida/eDot/pkg/http/handler"
	customMiddleware "github.com/fahmyabida/eDot/pkg/http/middleware"
)

var productCommand = &cobra.Command{
	Use:   "product",
	Short: "Start Product server",
	Run:   RunProductServer,
}

func init() {
	rootCmd.AddCommand(productCommand)
}

func RunProductServer(cmd *cobra.Command, args []string) {
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
	httpHandler.InitProductHandler(v1, ProductUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
