package districts

import "github.com/jinzhu/gorm"

type Districts struct {
  gorm.Model
  Name string
}

func (Districts) TableName() string {
    return "tm_districts"
}