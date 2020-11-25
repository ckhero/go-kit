package model

import (
	databases "base-demo/pkg/db/mysql"
)

type Goods struct {
	OrderId  uint64 `gorm:"column:orderId;"`       // 订单ID
	SkuId    uint64 `gorm:"column:skuId;"`       // SKU ID
	GoodsId  uint64 `gorm:"column:goodsId;"`                            // 商品ID
	Price    uint64 `gorm:"column:price;"`   // 商品单价（单位：分）
	VipPrice uint64 `gorm:"column:vipPrice;"` // 会员价（单位：分）
	Point    uint64 `gorm:"column:point;"`           // 积分（单位：分）
	Count    uint32 `gorm:"column:count;"`            // 商品数量
	databases.Time
}

func (Goods) TableName() string {
	return "order_goods"
}
