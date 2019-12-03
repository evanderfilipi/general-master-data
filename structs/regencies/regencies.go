package regencies

import "github.com/jinzhu/gorm"

type Regencies struct {
  gorm.Model
  Name string
}

func (Regencies) TableName() string {
    return "tm_regencies"
}