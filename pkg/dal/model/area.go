/**
 *@Description
 *@ClassName area
 *@Date 2020/11/25 10:17 上午
 *@Author ckhero
 */

package model

import (
	databases "base-demo/pkg/db/mysql"
	"time"
)

type Area struct {
	AreaID    int64       `gorm:"column:areaId;primary_key" json:"areaId"`
	Name      string    `gorm:"column:name" json:"name"`
	Principal string    `gorm:"column:principal" json:"principal"`
	Mobile    string    `gorm:"column:mobile" json:"mobile"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatetAt time.Time `gorm:"column:createtAt" json:"createtAt"`
	UpdateAt  time.Time `gorm:"column:updateAt" json:"updateAt"`
	databases.BasicTimeFields
}

// TableName sets the insert table name for this struct type
func (a *Area) TableName() string {
	return "area"
}
