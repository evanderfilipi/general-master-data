package provinces

import (
	// "net/http"
	// "fmt"
	// "time"
	"../../helper"
	"../../structs/provinces"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) GetDetail(c *gin.Context) {
	// fmt.Println(c)
	var (
		prov   provinces.Provinces
		res    helper.Response
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&prov).Error
	if err != nil {
		// res.Error = true
		// res.Code = 400
		// res.Message = err.Error()
		// res.Message = ""
		// res.Data = nil
		// c.JSON(http.StatusBadRequest, res)
		// 15:04:05
		// dn := time.Now()
		// dt := "2019-10-14"
		// // dt := "test123 test123"
		// hai := helper.DateToTimestamp(dt)
		// fmt.Println(hai)
		// dt2 := "1571011200"
		// hai2 := helper.TimestampToDate(dt2)
		// fmt.Println(hai2)
		// fmt.Println(dn)
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		result = gin.H{
			"provinces": prov,
		}
		res.Error = false
		res.Code = 200
		res.Message = "Success!"
		res.Data = result
		// c.JSON(http.StatusOK, res)
		helper.Responses(res, c)
	}
}

func (idb *InDB) GetList(c *gin.Context) {
	var (
		prov   []provinces.Provinces
		res    helper.Response
		result gin.H
	)

	idb.DB.Find(&prov)
	if len(prov) <= 0 {
		result = gin.H{
			"total_records": 0,
			"provinces":     nil,
		}
		res.Error = false
		res.Code = 200
		res.Message = "Success!"
		res.Data = result
	} else {
		result = gin.H{
			"total_records": len(prov),
			"provinces":     prov,
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
	data := make([]provinces.Provinces, 0)
	for _, value := range c.Request.PostForm {
		// fmt.Println(len(value))
		for i := 0; i < len(value); i++ {
			var (
				prov provinces.Provinces
			)
			prov.Name = value[i]
			idb.DB.Create(&prov)
			data = append(data, prov)
		}
		records = len(value)
	}
	result = gin.H{
		"total_records": records,
		"provinces":     data,
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
		prov    provinces.Provinces
		newProv provinces.Provinces
		res     helper.Response
		result  gin.H
	)

	err := idb.DB.Where("id = ?", id).First(&prov).Error
	if err != nil {
		// result = gin.H{
		// 	"success": false,
		// }
		// res.Error = true
		// res.Code = 400
		// res.Message = "Data tidak ditemukan!"
		// res.Data = result
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		newProv.Name = name
		err = idb.DB.Model(&prov).Updates(newProv).Error
		if err != nil {
			// result = gin.H{
			// 	"success": false,
			// }
			// res.Error = true
			// res.Code = 400
			// // res.Message = err.Error()
			// res.Message = "Data gagal diupdate!"
			// res.Data = result
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
		prov   provinces.Provinces
		res    helper.Response
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&prov).Error
	if err != nil {
		// result = gin.H{
		// 	"success": false,
		// }
		// res.Error = true
		// res.Code = 400
		// res.Message = "Data tidak ditemukan!"
		// res.Data = result
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		err = idb.DB.Delete(&prov).Error
		if err != nil {
			// result = gin.H{
			// 	"success": false,
			// }
			// res.Error = true
			// res.Code = 400
			// // res.Message = err.Error()
			// res.Message = "Data gagal dihapus!"
			// res.Data = result
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
