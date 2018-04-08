package gorm

import (
	jGorm "github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

func RegisterCallbacks(db *jGorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", beforeCreate)
	db.Callback().Update().Before("gorm:update").Register("my_plugin:before_update", beforeUpdate)
	db.Callback().Delete().Before("gorm:delete").Register("my_plugin:before_delete", beforeDelete)
	db.Callback().Query().Before("gorm:query").Register("my_plugin:before_query", beforeQuery)
}

func beforeQuery(scope *jGorm.Scope) {
	scope.Search.Where("deleted = 0")
}

func beforeCreate(scope *jGorm.Scope) {
	id := uuid.NewV4().String()

	field, ok := scope.FieldByName("id")
	if ok == true && field.IsBlank {
		err := scope.SetColumn("id", id)
		if err != nil {
			log.Fatal("err callback id ", err)
		}
	}

	timeNow := time.Now().Unix()

	field, ok = scope.FieldByName("created")
	if ok == true && field.IsBlank {
		err := scope.SetColumn("created", timeNow)
		if err != nil {
			log.Fatal("err callback created ", err)
		}
	}

	field, ok = scope.FieldByName("updated")
	if ok == true && field.IsBlank {
		err := scope.SetColumn("updated", timeNow)
		if err != nil {
			log.Fatal("err callback updated ", err)
		}
	}
}

func beforeUpdate(scope *jGorm.Scope) {
	field, ok := scope.FieldByName("updated")
	if ok == true && field.IsBlank {
		err := scope.SetColumn("updated", time.Now().Unix())
		if err != nil {
			log.Fatal("err callback updated ", err)
		}
	}
}

func beforeDelete(scope *jGorm.Scope) {
	field, ok := scope.FieldByName("deleted")
	if ok == true && field.IsBlank {
		err := scope.SetColumn("deleted", time.Now().Unix())
		if err != nil {
			log.Fatal("err callback deleted ", err)
		}
	}
}
