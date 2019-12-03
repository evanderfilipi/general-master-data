package main

import (
	"./config"
	"./controllers/provinces"
	"./controllers/regencies"
	"./controllers/districts"
	"./controllers/mapping"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	provTbl := &provinces.InDB{DB: db}
	regenTbl := &regencies.InDB{DB: db}
	disTbl := &districts.InDB{DB: db}
	mapTbl := &mapping.InDB{DB: db}

	// _ = regenTbl
	// _ = disTbl

	router := gin.Default()

	//provinces
	router.GET("/provinces/:id", provTbl.GetDetail)
	router.GET("/provinces", provTbl.GetList)
	router.POST("/provinces", provTbl.Create)
	router.PUT("/provinces", provTbl.Update)
	router.DELETE("/provinces/:id", provTbl.Delete)

	//regencies
	router.GET("/regencies/:id", regenTbl.GetDetail)
	router.GET("/regencies", regenTbl.GetList)
	router.POST("/regencies", regenTbl.Create)
	router.PUT("/regencies", regenTbl.Update)
	router.DELETE("/regencies/:id", regenTbl.Delete)

	//districts
	router.GET("/districts/:id", disTbl.GetDetail)
	router.GET("/districts", disTbl.GetList)
	router.POST("/districts", disTbl.Create)
	router.PUT("/districts", disTbl.Update)
	router.DELETE("/districts/:id", disTbl.Delete)

	//mapping
	router.GET("/mapping", mapTbl.GetList)
	router.POST("/mappings", mapTbl.GetListByFilter)
	router.GET("/mapping/:id", mapTbl.GetDetail)
	router.POST("/mapping", mapTbl.Create)
	router.PUT("/mapping", mapTbl.Update)
	router.DELETE("/mapping/:id", mapTbl.Delete)

	router.Run(":3000")
}