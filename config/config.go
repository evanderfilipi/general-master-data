package config

import (
	"../structs/provinces"
	"../structs/regencies"
	"../structs/districts"
	"../structs/mapping"
	"github.com/jinzhu/gorm"
	"fmt"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ksidev_master_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	if !db.HasTable(provinces.Provinces{}) {
		db.CreateTable(provinces.Provinces{})
		table_name := db.NewScope(provinces.Provinces{}).TableName()
		fmt.Println("Tabel",table_name,"berhasil dibuat!")
	}

	if !db.HasTable(regencies.Regencies{}) {
		db.CreateTable(regencies.Regencies{})
		table_name := db.NewScope(regencies.Regencies{}).TableName()
		fmt.Println("Tabel",table_name,"berhasil dibuat!")
	}

	if !db.HasTable(districts.Districts{}) {
		db.CreateTable(districts.Districts{})
		table_name := db.NewScope(districts.Districts{}).TableName()
		fmt.Println("Tabel",table_name,"berhasil dibuat!")
	}

	if !db.HasTable(mapping.Mapping{}) {
		db.CreateTable(mapping.Mapping{})
		table_name := db.NewScope(mapping.Mapping{}).TableName()
		fmt.Println("Tabel",table_name,"berhasil dibuat!")
	}

	// db.AutoMigrate(structs.Provinces{})
	return db
}