package database

import (
	"beluga/server/common/logger"
	"beluga/utils"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func FuzzyQuery(db *gorm.DB, instance interface{}, query string, ignores []string) *gorm.DB {
	if query == "" {
		return db
	}
	t := reflect.TypeOf(instance)
	v := reflect.ValueOf(instance)
	log := logger.GetLogger()
	log.Info("in fuzzy query, query: ", query)
	switch t.Kind() {
	case reflect.Struct:
		log.Info("structure case...")
		for i := 0; i < t.NumField(); i++ {
			isIgnore := false
			for _, ignore := range ignores {
				if t.Field(i).Name == ignore {
					isIgnore = true
					break
				}
			}
			if isIgnore {
				continue
			}
			value := v.Field(i)
			switch value.Kind() {
			case reflect.String:
				{
					condition := fmt.Sprintf("%s LIKE '%%%s%%'", utils.Camel2Case(t.Field(i).Name), query)
					db = db.Or(condition)
				}
			}
			if db.Error != nil {
				return db
			}
		}
	default:
		log.Info("default case...")
		return db
	}
	return db
}
