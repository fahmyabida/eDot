package postgres

// import (
// 	"strings"

// 	migrate "github.com/golang-migrate/migrate/v4"
// 	_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
// 	"gorm.io/gorm"
// )

// type Migration struct {
// 	Migrate *migrate.Migrate
// }

// func (m *Migration) Up() (bool, error) {
// 	err := m.Migrate.Up()
// 	if err != nil {
// 		if err == migrate.ErrNoChange {
// 			return true, nil
// 		}
// 		return false, err
// 	}
// 	return true, nil
// }

// func (m *Migration) Down() (bool, error) {
// 	err := m.Migrate.Down()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, err
// }

// func RunMigration(dbConn *gorm.DB, migrationsFolderLocation string) (*Migration, error) {
// 	dataPath := []string{}
// 	dataPath = append(dataPath, "file://")
// 	dataPath = append(dataPath, migrationsFolderLocation)

// 	pathToMigrate := strings.Join(dataPath, "")
// 	db, _ := dbConn.DB()
// 	driver, err := _postgres.WithInstance(db, &_postgres.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, postgres, driver)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Migration{Migrate: m}, nil
// }
