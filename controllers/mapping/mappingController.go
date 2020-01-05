package mapping

import (
	// "net/http"
	"fmt"
	"strings"

	// "io/ioutil"
	// "../../structs/provinces"
	// "../../structs/regencies"
	// "../../structs/districts"
	"../../handler"
	"../../helper"
	"../../structs/mapping"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "encoding/json"
	// "bytes"
)

type InDB struct {
	DB *gorm.DB
}

type Set struct {
	Next      int64       `json:"next,omitempty"`
	Records   int         `json:"total_records,omitempty"`
	Provinces interface{} `json:"provinces,omitempty"`
	Results   interface{} `json:"results,omitempty"`
}

func (idb *InDB) GetDetail(c *gin.Context) {
	var (
		res helper.Response
		// set Set
		// records int
	)

	id := c.Param("id")
	selectItems := "tn_provinceregencycitydistricts.id, tm_provinces.id, tm_provinces.name, tm_regencies.id, tm_regencies.name, tm_districts.id, tm_districts.name"
	db := idb.DB.Table("tn_provinceregencycitydistricts").
		Joins("JOIN tm_provinces ON tn_provinceregencycitydistricts.province_id = tm_provinces.id").
		Joins("JOIN tm_regencies ON tn_provinceregencycitydistricts.regency_city_id = tm_regencies.id").
		Joins("JOIN tm_districts ON tn_provinceregencycitydistricts.district_id = tm_districts.id").
		Where("tn_provinceregencycitydistricts.id=?", id)

	data, err := db.Select(selectItems).Rows()
	if err != nil {
		msg := err.Error()
		helper.ErrorCustomStatus(400, msg, c)
	}
	defer data.Close()

	datas := []mapping.Mapped{}
	// dis := []mapping.District{}
	// var maps mapping.Mapped

	for data.Next() {
		var mapped mapping.Mapped
		err := data.Scan(&mapped.MappingId, &mapped.ProvinceId, &mapped.ProvinceName, &mapped.RegencyId, &mapped.RegencyName, &mapped.DistrictId, &mapped.DistrictName)
		if err != nil {
			msg := err.Error()
			helper.ErrorCustomStatus(400, msg, c)
		}
		datas = append(datas, mapped)
	}
	// set.Results = datas

	res.Error = false
	res.Code = 200
	res.Message = "Success!"
	res.Data = datas

	helper.Responses(res, c)
}

func (idb *InDB) GetList(c *gin.Context) {
	var (
		res     helper.Response
		set     Set
		records int
	)

	// selectItems := "tm_provinces.id, tm_provinces.name, tm_regencies.id, tm_regencies.name, tm_districts.id, tm_districts.name"
	db := idb.DB.Table("tn_provinceregencycitydistricts").
		Joins("JOIN tm_provinces ON tn_provinceregencycitydistricts.province_id = tm_provinces.id").
		Joins("JOIN tm_regencies ON tn_provinceregencycitydistricts.regency_city_id = tm_regencies.id").
		Joins("JOIN tm_districts ON tn_provinceregencycitydistricts.district_id = tm_districts.id")

	// prov, err := db.Select(selectItems).Rows()
	// if err != nil {
	// 	msg := err.Error()
	//     helper.ErrorCustomStatus(400, msg, c)
	// }
	// defer prov.Close()

	// group1 := make(map[int64][]mapping.Mapped)
	// group2 := make(map[int64][]mapping.Mapped)
	// group3 := make(map[int64][]mapping.Mapped)
	// datas := []mapping.Mapped{}
	// // dis := []mapping.District{}
	// // var maps mapping.Mapped

	// for prov.Next() {
	//     var mapped mapping.Mapped
	//     err := prov.Scan(&mapped.ProvinceId, &mapped.ProvinceName, &mapped.RegencyId, &mapped.RegencyName, &mapped.DistrictId, &mapped.DistrictName)
	//     if err != nil {
	//         msg := err.Error()
	//     	helper.ErrorCustomStatus(400, msg, c)
	//     }

	// 	group1[mapped.ProvinceId] = append(group1[mapped.ProvinceId], mapped)
	// 	group2[mapped.RegencyId] = append(group2[mapped.RegencyId], mapped)
	// 	group3[mapped.DistrictId] = append(group3[mapped.DistrictId], mapped)
	// 	datas = append(datas, mapped)

	// }

	items1 := "tm_provinces.id, tm_provinces.name"

	data1, err := db.Select(items1).Group("tm_provinces.id").Rows()
	if err != nil {
		msg := err.Error()
		helper.ErrorCustomStatus(400, msg, c)
	}
	defer data1.Close()
	list1 := []mapping.Province{}

	for data1.Next() {
		var map1 mapping.Province
		err := data1.Scan(&map1.Id, &map1.Name)
		if err != nil {
			msg := err.Error()
			helper.ErrorCustomStatus(400, msg, c)
		}
		list1 = append(list1, map1)
	}

	items2 := "tm_regencies.id, tm_regencies.name, tm_provinces.id"

	data2, err := db.Select(items2).Group("tm_regencies.id").Rows()
	if err != nil {
		msg := err.Error()
		helper.ErrorCustomStatus(400, msg, c)
	}
	defer data2.Close()
	list2 := []mapping.Regency{}

	for data2.Next() {
		var map2 mapping.Regency
		err := data2.Scan(&map2.Id, &map2.Name, &map2.ProvinceId)
		if err != nil {
			msg := err.Error()
			helper.ErrorCustomStatus(400, msg, c)
		}
		list2 = append(list2, map2)
	}

	items3 := "tm_districts.id, tm_districts.name, tm_regencies.id"

	data3, err := db.Select(items3).Group("tm_districts.id").Rows()
	if err != nil {
		msg := err.Error()
		helper.ErrorCustomStatus(400, msg, c)
	}
	defer data3.Close()
	list3 := []mapping.District{}

	for data3.Next() {
		var map3 mapping.District
		err := data3.Scan(&map3.Id, &map3.Name, &map3.RegencyId)
		if err != nil {
			msg := err.Error()
			helper.ErrorCustomStatus(400, msg, c)
		}
		list3 = append(list3, map3)
	}
	records = len(list3)

	get1 := []mapping.Province{}

	for _, value1 := range list1 {
		var m1 mapping.Province
		m1.Id = value1.Id
		m1.Name = value1.Name
		get2 := []mapping.Regency{}
		for _, value2 := range list2 {
			if value2.ProvinceId == value1.Id {
				var m2 mapping.Regency
				m2.Id = value2.Id
				m2.Name = value2.Name
				// m2.ProvinceId = value2.ProvinceId
				get3 := []mapping.District{}
				for _, value3 := range list3 {
					if value3.RegencyId == value2.Id {
						var m3 mapping.District
						m3.Id = value3.Id
						m3.Name = value3.Name
						// m3.RegencyId = value3.RegencyId
						get3 = append(get3, m3)
					}
				}
				m2.Dis = get3
				get2 = append(get2, m2)
			}
		}
		m1.Reg = get2
		get1 = append(get1, m1)
	}

	set.Records = records
	set.Provinces = get1

	res.Error = false
	res.Code = 200
	res.Message = "Success!"
	res.Data = set

	helper.Responses(res, c)
}

func (idb *InDB) GetListByFilter(c *gin.Context) {
	var (
		fils        mapping.Filtering
		res         helper.Response
		set         Set
		records     int
		selectItems string
	)

	// id := c.Param("id")
	c.ShouldBindJSON(&fils)

	// selectItems := "tn_provinceregencycitydistricts.id, tm_provinces.id, tm_provinces.name, tm_regencies.id, tm_regencies.name, tm_districts.id, tm_districts.name"
	db := idb.DB.Table("tn_provinceregencycitydistricts").
		Joins("JOIN tm_provinces ON tn_provinceregencycitydistricts.province_id = tm_provinces.id").
		Joins("JOIN tm_regencies ON tn_provinceregencycitydistricts.regency_city_id = tm_regencies.id").
		Joins("JOIN tm_districts ON tn_provinceregencycitydistricts.district_id = tm_districts.id")

	searchable := [3]string{"tm_provinces.name", "tm_regencies.name", "tm_districts.name"}

	if fils.Fields != nil {
		selectItems = strings.Join(fils.Fields, ",")
	} else {
		selectItems = "tn_provinceregencycitydistricts.id as mapping_id, tm_provinces.id as province_id, tm_provinces.name as province_name, tm_regencies.id as regency_city_id, tm_regencies.name as regency_city_name, tm_districts.id as dictrict_id, tm_districts.name as dictrict_name"
	}

	if fils.Keyword != "" {
		// for sc := 0; sc < len(searchable); sc++ {
		db = db.Where(searchable[0]+" LIKE ? OR "+searchable[1]+" LIKE ? OR "+searchable[2]+" LIKE ?", "%"+fils.Keyword+"%", "%"+fils.Keyword+"%", "%"+fils.Keyword+"%")
		// }
	}

	if fils.Filter != nil {
		for k := 0; k < len(fils.Filter); k++ {
			db = db.Where(fils.Filter[k].Key+" = ?", fils.Filter[k].Value)
		}
	}

	if fils.ArrayFilter != nil {
		for m := 0; m < len(fils.ArrayFilter); m++ {
			db = db.Where(fils.ArrayFilter[m].Key+" IN (?)", fils.ArrayFilter[m].Value)
		}
	}

	if fils.DateFilter != nil {
		for d := 0; d < len(fils.DateFilter); d++ {
			starts := helper.TimestampToDate(fils.DateFilter[d].Value.Start)
			ends := helper.TimestampToDate(fils.DateFilter[d].Value.End)
			db = db.Where("tn_provinceregencycitydistricts."+fils.DateFilter[d].Key+" BETWEEN ? AND ?", starts, ends)
		}
	}

	if fils.Sort != "" && fils.SortBy != "" {
		db = db.Order(fils.Sort + " " + fils.SortBy)
	}

	db = db.Count(&records)

	if fils.Page != 0 && fils.PerPage != 0 {
		db = db.Offset((fils.Page - 1) * fils.PerPage).Limit(fils.PerPage)
		set.Next = fils.Page + 1
	}

	fmt.Println(selectItems)
	data, err := db.Select(selectItems).Rows()
	if err != nil {
		msg := err.Error()
		helper.ErrorCustomStatus(400, msg, c)
	}

	columns, errs := data.Columns()
	if errs != nil {
		msg := errs.Error()
		helper.ErrorCustomStatus(400, msg, c)
	}
	length := len(columns)
	result := make([]map[string]interface{}, 0)

	defer data.Close()

	// datas := []mapping.Mapped{}
	// dis := []mapping.District{}
	// var maps mapping.Mapped

	for data.Next() {
		// var mapped mapping.Mapped
		// err := data.Scan(&mapped.MappingId, &mapped.ProvinceId, &mapped.ProvinceName, &mapped.RegencyId, &mapped.RegencyName, &mapped.DistrictId, &mapped.DistrictName)
		// if err != nil {
		// 	msg := err.Error()
		// 	helper.ErrorCustomStatus(400, msg, c)
		// }
		// datas = append(datas, mapped)
		current := handler.MakeResultReceiver(length)
		if err := data.Scan(current...); err != nil {
			msg := err.Error()
			helper.ErrorCustomStatus(400, msg, c)
		}
		value := handler.ScanResult(length, current, columns)
		result = append(result, value)
	}
	set.Records = records
	set.Results = result

	res.Error = false
	res.Code = 200
	res.Message = "Success!"
	res.Data = set

	helper.Responses(res, c)
}

func (idb *InDB) Create(c *gin.Context) {
	var (
		maps mapping.MappingData
		res  helper.Response
		set  Set
	)
	data := make([]mapping.InputData, 0)
	c.ShouldBindJSON(&maps)
	for i := 0; i < len(maps); i++ {
		var (
			mp mapping.InputData
		)
		mp.ProvinceId = maps[i].ProvinceID
		mp.RegencyCityId = maps[i].RegencyID
		mp.DistrictId = maps[i].DistrictID
		idb.DB.Create(&mp)
		data = append(data, mp)
	}
	set.Records = len(data)
	set.Results = data

	res.Error = false
	res.Code = 200
	res.Message = "Success!"
	res.Data = set

	helper.Responses(res, c)
	// c.JSON(http.StatusOK, maps)
	// fmt.Println(maps[0].ProvinceID)
	// var (
	// 	set Set
	// 	res helper.Response
	// 	records int
	// )
	// maps := make(map[string]interface{})
	// err := json.NewDecoder(c.Request.Body).Decode(&maps)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(maps)
	// body := c.Request.Body
	// bodyBytes, _ = ioutil.ReadAll(body)
	// body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// // maps := new(mapping.Final)
	// err := c.Bind(&maps)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(maps)
	// c.Request.ParseMultipartForm(1000)
	// c.Request.Form.Get("data")
	// mydata := c.PostForm("data")
	// fmt.Println(mydata)
	// for _, val := range maps {
	// 	fmt.Println(val)
	// }
	// fmt.Printf("%s \n", string(x))

	// c.Request.ParseMultipartForm(1000)
	// data := c.PostFormArray("data")
	// data := make([]provinces.Provinces, 0)
	// for _, value := range c.Request.PostForm {
	// 	fmt.Println(value)
	// for _, v := range value {
	// 	// Do something with the i-th value for key k.
	// }
	// for i := 0; i < len(value); i++ {
	// 	var (
	// 		prov provinces.Provinces
	// 	)
	// 	prov.Name = value[i]
	// 	idb.DB.Create(&prov)
	// 	data = append(data, prov)
	// }
	// records = len(value)
	// }
	// result = gin.H {
	// 	"total_records" : records,
	// 	"provinces" : data,
	// }
	// res.Error = false
	// res.Code = 200
	// res.Message = "Data berhasil disimpan!"
	// res.Data = result
	// helper.Responses(res, c)
}

func (idb *InDB) Update(c *gin.Context) {
	// id := c.Query("id")
	var (
		maps mapping.MappingData
		res  helper.Response
		set  Set
	)
	data := make([]mapping.InputData, 0)
	c.ShouldBindJSON(&maps)

	for i := 0; i < len(maps); i++ {
		var (
			mp    mapping.InputData
			newmp mapping.InputData
		)
		err := idb.DB.Where("id = ?", maps[i].MappingID).First(&mp).Error
		if err == nil {
			newmp.ProvinceId = maps[i].ProvinceID
			newmp.RegencyCityId = maps[i].RegencyID
			newmp.DistrictId = maps[i].DistrictID
			idb.DB.Model(&mp).Updates(newmp)
			data = append(data, newmp)
		}
	}

	set.Records = len(data)
	set.Results = data

	res.Error = false
	res.Code = 200
	res.Message = "Success!"
	res.Data = set

	helper.Responses(res, c)

	// if err != nil {
	// 	// result = gin.H{
	// 	// 	"success": false,
	// 	// }
	// 	// res.Error = true
	// 	// res.Code = 400
	// 	// res.Message = "Data tidak ditemukan!"
	// 	// res.Data = result
	// 	helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	// } else {
	// 	newProv.Name = name
	// 	err = idb.DB.Model(&prov).Updates(newProv).Error
	// 	if err != nil {
	// 		// result = gin.H{
	// 		// 	"success": false,
	// 		// }
	// 		// res.Error = true
	// 		// res.Code = 400
	// 		// // res.Message = err.Error()
	// 		// res.Message = "Data gagal diupdate!"
	// 		// res.Data = result
	// 		helper.ErrorCustomStatus(400, "Data gagal diupdate!", c)
	// 	} else {
	// 		result = gin.H{
	// 			"success": true,
	// 		}
	// 		res.Error = false
	// 		res.Code = 200
	// 		res.Message = "Data berhasil diupdate!"
	// 		res.Data = result
	// 		helper.Responses(res, c)
	// 	}
	// }
}

func (idb *InDB) Delete(c *gin.Context) {
	var (
		mp  mapping.InputData
		res helper.Response
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&mp).Error
	if err != nil {
		helper.ErrorCustomStatus(400, "Data tidak ditemukan!", c)
	} else {
		err = idb.DB.Delete(&mp).Error
		if err != nil {
			helper.ErrorCustomStatus(400, "Data gagal dihapus!", c)
		} else {
			res.Error = false
			res.Code = 200
			res.Message = "Success!"

			helper.Responses(res, c)
		}
	}
}

// func (s helper.Set) MarshalJSON() ([]byte, error) {
// 	key := "provinces"
//     data := map[string]interface{}{
//         key: s.Results,
//     }
//     return json.Marshal(data)
// }
