package mapping

import (
	"github.com/jinzhu/gorm"
	// "../../structs/provinces"
	// "../../structs/regencies"
)

type Mapping struct {
	gorm.Model
	ProvinceId    int
	RegencyCityId int
	DistrictId    int
}

type MappingData []struct {
	// Data []struct{
	MappingID  int64 `form:"mapping_id" json:"mapping_id,omitempty"`
	ProvinceID int64 `form:"province_id" json:"province_id,omitempty"`
	RegencyID  int64 `form:"regency_id" json:"regency_id,omitempty"`
	DistrictID int64 `form:"district_id" json:"district_id,omitempty"`
	// } `form:"data" json:"data"`
}

type Filtering struct {
	Fields     []string `form:"fields" json:"fields,omitempty"`
	Keyword    string   `form:"keyword" json:"keyword,omitempty"`
	DateFilter []struct {
		Key   string `form:"key" json:"key,omitempty"`
		Value struct {
			Start int64 `form:"start" json:"start,omitempty"`
			End   int64 `form:"end" json:"end,omitempty"`
		} `form:"value" json:"value,omitempty"`
	} `form:"date_filter" json:"date_filter,omitempty"`
	Page    int64 `form:"page" json:"page,omitempty"`
	PerPage int64 `form:"per_page" json:"per_page,omitempty"`
	// Filter struct {
	// 	Key   []string `form:"key" json:"key,omitempty"`
	// 	Value []string `form:"value" json:"value,omitempty"`
	// } `form:"filter" json:"filter,omitempty"`
	Filter []struct {
		Key   string `form:"key" json:"key,omitempty"`
		Value string `form:"value" json:"value,omitempty"`
	} `form:"filter" json:"filter,omitempty"`
	ArrayFilter []struct {
		Key   string   `form:"key" json:"key,omitempty"`
		Value []string `form:"value" json:"value,omitempty"`
	} `form:"array_filter" json:"array_filter,omitempty"`
	Sort   string `form:"sort" json:"sort,omitempty"`
	SortBy string `form:"sort_by" json:"sort_by,omitempty"`
}

type Mapped struct {
	MappingId    int64  `json:",omitempty"`
	ProvinceId   int64  `json:",omitempty"`
	ProvinceName string `json:",omitempty"`
	RegencyId    int64  `json:",omitempty"`
	RegencyName  string `json:",omitempty"`
	DistrictId   int64  `json:",omitempty"`
	DistrictName string `json:",omitempty"`
}

type InputData struct {
	gorm.Model
	ProvinceId    int64 `json:"province_id"`
	RegencyCityId int64 `json:"regency_id"`
	DistrictId    int64 `json:"district_id"`
}

func (InputData) TableName() string {
	return "tn_provinceregencycitydistricts"
}

// type Province struct {
//     // Prov provinces.Provinces
//     Id int64
//     Name string
//     // Reg Regency `json:"regencies,omitempty"`
// }

// type Regency struct {
//     Id int64
//     Name string
//     // Dis District `json:"district,omitempty"`
// }

// type District struct {
//     Id int64
//     Name string
// }

type Province struct {
	Id   int64
	Name string
	Reg  []Regency `json:"regencies,omitempty"`
}

type Regency struct {
	Id         int64
	Name       string
	ProvinceId int64      `json:",omitempty"`
	Dis        []District `json:"districts,omitempty"`
}

type District struct {
	Id        int64
	Name      string
	RegencyId int64 `json:",omitempty"`
}

func (Mapping) TableName() string {
	return "tn_provinceregencycitydistricts"
}
