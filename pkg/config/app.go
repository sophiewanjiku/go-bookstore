package config

import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/dialects/mysql"
)

var(
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "sophie:qwerty/simplerest?charset=utf8&parseTime=True&loc=loclatime")
	if err != nil{
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}