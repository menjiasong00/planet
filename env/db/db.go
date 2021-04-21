package env

import (
	"github.com/jinzhu/gorm"
	"planet/env"
	"planet/pkg/gmysql"
)

var SelectDB  = make(map[string]*gorm.DB,0)

func New(name string) *gorm.DB {
	if db, ok := SelectDB[name]; ok {
		return db
	}else{
		var ConnConfig gmysql.ConnConfig
		env.ScanConfig("Mysql."+name,&ConnConfig)
		oneDB := gmysql.New(ConnConfig)
		SelectDB[name] = oneDB
		return oneDB
	}
}