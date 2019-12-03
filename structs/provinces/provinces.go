package provinces

import "github.com/jinzhu/gorm"

type Provinces struct {
  gorm.Model
  Name string
}

func (Provinces) TableName() string {
    return "tm_provinces"
}