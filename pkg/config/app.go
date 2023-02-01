package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func connect() {
	d, err := gorm.Open("mysql", "sophie:qwerty/simplerest?")
	if err != nil {
		panic(err)
	}
	db = d

}
