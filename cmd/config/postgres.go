package config

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Postgres struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Dbname   string `envconfig:"DATABASE" required:"true"`

	MaxConnectionLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFE_TIME" default:"300s"`
	MaxOpenConnection     int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"100"`
	MaxIdleConnection     int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
}

func (p Postgres) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

// InitDB initializes the database connection
func InitDB(rw, ro *Postgres) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
			rw.Host,
			rw.Port,
			rw.User,
			rw.Password,
			rw.Dbname,
			"disable",
		),
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect rw database")
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(rw.MaxConnectionLifetime)
	sqlDB.SetMaxOpenConns(rw.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(rw.MaxIdleConnection)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		panic(err)
	}

	if ro != nil {
		if err := db.Use(readOnlyResolver(ro)); err != nil {
			panic(err)
		}
	}

	return db
}

func readOnlyResolver(ro *Postgres) *dbresolver.DBResolver {
	dbReadOnlyConn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
			ro.Host,
			ro.Port,
			ro.User,
			ro.Password,
			ro.Dbname,
			"disable",
		),
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect ro database")
	}

	resolver := dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.New(postgres.Config{Conn: dbReadOnlyConn.ConnPool})},
	})
	resolver.SetConnMaxLifetime(ro.MaxConnectionLifetime)
	resolver.SetMaxOpenConns(ro.MaxOpenConnection)
	resolver.SetMaxIdleConns(ro.MaxIdleConnection)

	return resolver
}

// LoadForPostgres loads postgres configuration and returns it
func LoadForPostgres() (rw *Postgres, ro *Postgres) {
	rw = &Postgres{}
	ro = &Postgres{}

	mustLoad("PG", rw)

	if err := envconfig.Process("PG_RO", ro); err != nil {
		log.Default().Print(
			"Failed to load read-only database configuration, using the read-write database.",
		)

		return rw, nil
	}

	return
}
