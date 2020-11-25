/**
 *@Description
 *@ClassName goods
 *@Date 2020/11/25 9:30 上午
 *@Author ckhero
 */

package dao

import (
	order "base-demo/pkg/dal/model"
	"base-demo/pkg/db/mysql"
	"context"
)

type goods struct {
	mysql.Base
}

func NewGoodsDao() *goods {
	return &goods{}
}

func (g *goods) First(ctx context.Context) (*order.Goods, error) {
	out := order.Goods{}
	db := g.GetDefaultDB(ctx)
	err := db.First(out).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}