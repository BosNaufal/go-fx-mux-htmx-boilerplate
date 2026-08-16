package main

import (
	"fmt"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	infra "gitlab.com/bosnaufal/bos-ai-search/internal/infrastructures"
)

func MigrateSeed(dbmateInstance *dbmate.DB) error {
	dbmateInstance.MigrationsDir = []string{"./migrations/seeder"}

	fmt.Println("running migration init")
	err := dbmateInstance.CreateAndMigrate()
	if err != nil {
		panic(err)
	}
	return err
}

func MigrateUp(dbmateInstance *dbmate.DB) error {
	dbmateInstance.MigrationsDir = []string{"./migrations/"}

	fmt.Println("running migration up")
	err := dbmateInstance.CreateAndMigrate()
	if err != nil {
		panic(err)
	}
	return err
}

func MigrateDown(dbmateInstance *dbmate.DB) error {
	dbmateInstance.MigrationsDir = []string{"./migrations/"}

	fmt.Println("rollback migration")
	err := dbmateInstance.Rollback()
	if err != nil {
		panic(err)
	}

	return err
}

func main() {
	fmt.Printf("%#v", os.Args)
	if len(os.Args) < 2 {
		panic("need one argument 'up' or 'down'")
	}

	var err error

	migrationType := os.Args[1]

	fmt.Println("migration type", migrationType)

	dbmateInstance := infra.GetSQLStoreMigration(
		"string",
		"string",
		"string",
		"string",
		"string",
	)

	switch os.Args[1] {
	case "seed":
		fmt.Println("migrate init")
		err = MigrateSeed(dbmateInstance)
	case "up":
		fmt.Println("migrate up")
		err = MigrateUp(dbmateInstance)
	case "down":
		fmt.Println("migrate down")
		err = MigrateDown(dbmateInstance)
	}

	if err != nil {
		fmt.Println("%#v", err)
	}
}
