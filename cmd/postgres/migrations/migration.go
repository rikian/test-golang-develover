package migration

import (
	"go/service1/config"
	t "go/service1/shared/models/entities/table"
	"log"
)

type table struct {
	tableName interface{}
}

func registerTable() []*table {
	return []*table{
		{tableName: &t.StatusUser{}},
		{tableName: &t.Categories{}},
		{tableName: &t.User{}},
		{tableName: &t.Product{}},
	}
}

func RunMigration() {
	config.LoadEnvFile()

	db := *config.ConnectDB()

	for _, table := range registerTable() {
		dropTable := db.Migrator().DropTable(table.tableName)

		if dropTable != nil {
			log.Print(dropTable.Error())
		}
	}

	err := db.Migrator().AutoMigrate(
		&t.StatusUser{},
		&t.User{},
		&t.Categories{},
		&t.Product{},
	)

	if err != nil {
		log.Print(err.Error())
	}

	err = db.Migrator().CreateConstraint(&t.User{}, "Products")

	if err != nil {
		log.Print(err.Error())
	}

	if !db.Migrator().HasConstraint(&t.User{}, "Products") {
		log.Print("Failed create Constraint at tb_user")
	}

	if !db.Migrator().HasConstraint(&t.User{}, "fk_public_users_products") {
		log.Print("Failed create Constraint at tb_user")
	}

	tbStatus := db.Create(&t.StatusUser{
		Id:     1,
		Status: "admin",
	})

	if tbStatus.Error != nil {
		log.Print(tbStatus.Error.Error())
	}

	tbCategories := db.Create(&t.Categories{
		Id:   "c-001kntl",
		Name: "electronic",
	})

	if tbCategories.Error != nil {
		log.Print(tbCategories.Error.Error())
	}

	log.Print("success migration table")
}
