package database

import "beluga/utils"

func MigrateDatabase(generators []utils.Generator) error {
	db := GetDB()
	for _, g := range generators {
		db.AutoMigrate(g(utils.PTR))
	}
	return nil
}
