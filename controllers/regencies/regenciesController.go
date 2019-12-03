package regencies

import (

	// "fmt"
	"../../helper"
	"../../structs/regencies"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) GetDetail(c *gin.Context) {
	var (
		regen  regencies.Regencies
		res    helper.Response
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&regen).Error
	if err != nil {
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		result = gin.H{
			"regencies": regen,
		}
		res.Error = false
		res.Code = 200
		res.Message = "Success!"
		res.Data = result
		helper.Responses(res, c)
	}
}

func (idb *InDB) GetList(c *gin.Context) {
	var (
		regen  []regencies.Regencies
		res    helper.Response
		result gin.H
	)

	idb.DB.Find(&regen)
	if len(regen) <= 0 {
		result = gin.H{
			"total_records": 0,
			"regencies":     nil,
		}
		res.Error = false
		res.Code = 200
		res.Message = "Success!"
		res.Data = result
	} else {
		result = gin.H{
			"total_records": len(regen),
			"regencies":     regen,
		}
		res.Error = false
		res.Code = 200
		res.Message = "Success!"
		res.Data = result
	}

	helper.Responses(res, c)
}

func (idb *InDB) Create(c *gin.Context) {
	var (
		result  gin.H
		res     helper.Response
		records int
	)
	c.Request.ParseMultipartForm(1000)
	data := make([]regencies.Regencies, 0)
	for _, value := range c.Request.PostForm {
		for i := 0; i < len(value); i++ {
			var (
				regen regencies.Regencies
			)
			regen.Name = value[i]
			idb.DB.Create(&regen)
			data = append(data, regen)
		}
		records = len(value)
	}
	result = gin.H{
		"total_records": records,
		"regencies":     data,
	}
	res.Error = false
	res.Code = 200
	res.Message = "Data berhasil disimpan!"
	res.Data = result
	helper.Responses(res, c)
}

func (idb *InDB) Update(c *gin.Context) {
	// id := c.Query("id")
	id := c.PostForm("id")
	name := c.PostForm("name")
	var (
		regen    regencies.Regencies
		newRegen regencies.Regencies
		res      helper.Response
		result   gin.H
	)

	err := idb.DB.Where("id = ?", id).First(&regen).Error
	if err != nil {
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		newRegen.Name = name
		err = idb.DB.Model(&regen).Updates(newRegen).Error
		if err != nil {
			helper.ErrorCustomStatus(400, "Data gagal diupdate!", c)
		} else {
			result = gin.H{
				"success": true,
			}
			res.Error = false
			res.Code = 200
			res.Message = "Data berhasil diupdate!"
			res.Data = result
			helper.Responses(res, c)
		}
	}
}

func (idb *InDB) Delete(c *gin.Context) {
	var (
		regen  regencies.Regencies
		res    helper.Response
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&regen).Error
	if err != nil {
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		err = idb.DB.Delete(&regen).Error
		if err != nil {
			helper.ErrorCustomStatus(400, "Data gagal dihapus!", c)
		} else {
			result = gin.H{
				"success": true,
			}
			res.Error = false
			res.Code = 200
			res.Message = "Data berhasil dihapus!"
			res.Data = result
			helper.Responses(res, c)
		}
	}
}
