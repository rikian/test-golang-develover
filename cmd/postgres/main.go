package main

import migration "go/service1/cmd/postgres/migrations"

func main() {
	migration.RunMigration()
}
