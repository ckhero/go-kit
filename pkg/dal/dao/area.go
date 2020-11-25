/**
 *@Description
 *@ClassName area
 *@Date 2020/11/25 10:18 上午
 *@Author ckhero
 */

package dao

import (
	order "base-demo/pkg/dal/model"
	"base-demo/pkg/db/mysql"
	"context"
)

type area struct {
	mysql.Base
}

func NewAreaDao() *area {
	return &area{}
}

func (g *area) First(ctx context.Context) (*order.Area, error) {
	out := order.Area{}
	db := g.GetDefaultDB(ctx)
	err := db.Unscoped().First(out).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (g *area) Create(ctx context.Context) error {
	g.GetDefaultDB(ctx).Select("Name", "Principal", "Mobile", "Status").Create([]order.Area{{
		Name:      "测试",
		Principal: "郑杰",
		Mobile:    "13907316763",
		Status:    "VALID",
	},{
		Name:      "测试",
		Principal: "郑杰",
		Mobile:    "13907316763",
		Status:    "VALID",
	}})
	return nil
}