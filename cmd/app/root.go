package app

import (
	"log"

	"github.com/fahmyabida/eDot/cmd/config"
	"github.com/fahmyabida/eDot/internal/app/domain"
	"github.com/fahmyabida/eDot/internal/app/repository"
	"github.com/fahmyabida/eDot/internal/app/usecase"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	rootCmd = &cobra.Command{
		Use:   "brick-transfer-service",
		Short: "brick-transfer-service is application for transfer",
	}
)

var (
	// database
	database *gorm.DB

	// config
	jwtConfig *config.JWTConfig

	// repository
	UserRepo           domain.IUserRepo
	ProductRepo        domain.IProductRepo
	OrderRepo          domain.IOrderRepo
	OrderItemRepo      domain.IOrderItemRepo
	WarehouseRepo      domain.IWarehouseRepo
	WarehouseStockRepo domain.IWarehouseStockRepo

	// usecase
	UserUsecase      domain.IUserUsecase
	ProductUsecase   domain.IProductUsecase
	OrderUsecase     domain.IOrderUsecase
	WarehouseUsecase domain.IWarehouseUsecase
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	if err := config.InitEnv(); err != nil {
		log.Fatal(err)
	}

	cobra.OnInitialize(func() {
		initJWT()
		initDatabase()
		initApp()
	})
}

func initJWT() {
	jwtConfig = config.LoadForJWTConfig()
}

func initDatabase() {
	rw, ro := config.LoadForPostgres()
	database = config.InitDB(rw, ro)
}

func initApp() {

	UserRepo = repository.NewUserRepository(database)
	ProductRepo = repository.NewProductRepository(database)
	OrderRepo = repository.NewOrderRepository(database)
	OrderItemRepo = repository.NewOrderItemRepository(database)
	WarehouseRepo = repository.NewWarehouseRepository(database)
	WarehouseStockRepo = repository.NewWarehouseStockRepository(database)

	UserUsecase = usecase.NewUserUsecase(UserRepo, jwtConfig)
	ProductUsecase = usecase.NewProductUsecase(ProductRepo)
	OrderUsecase = usecase.NewOrderUsecase(ProductRepo, OrderRepo, OrderItemRepo)
	WarehouseUsecase = usecase.NewWarehouseUsecase(ProductRepo, WarehouseRepo, WarehouseStockRepo)
}
