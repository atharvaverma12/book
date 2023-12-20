package config

import {
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialect/mysql"
}

var {
	db * gorm.DB
}

func Connect(){
	d, err := gorm.Open("mysql", "atharva:12345678/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() * gorm.DB{
	return db
}