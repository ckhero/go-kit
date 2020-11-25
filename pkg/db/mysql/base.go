package mysql

import (
	"base-demo/pkg/constant"
	"context"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	db *gorm.DB
}

type Time struct {
	BasicTimeFields
	DeletedAt *time.Time `gorm:"column:deletedAt;null"`
}

type BasicTimeFields struct {
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}


/**
 * 获取数据库连接
 */
func (b *Base) GetDB(ctx context.Context, name ...string) *gorm.DB {
	if b.db == nil {
		b.SetDB(ctx, getDB(name...))
	}
	return b.db
}

func (b *Base) GetDefaultDB(ctx context.Context) *gorm.DB {
	return b.GetDB(ctx, constant.MysqlConfigKeyDefault)
}

/**
 * 设置数据库连接（仅当使用事物需要变更操作句柄时使用）
 */
func (b *Base) SetDB(ctx context.Context, db *gorm.DB) {
	b.db = SetSpanToGorm(ctx, db)
}


