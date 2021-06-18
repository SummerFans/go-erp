package entity

import "time"

// Goods ..
type Goods struct {
	GoodsID    int       `json:"goods_id"`
	CategoryID int       `json:"category_id"`
	Status     int       `json:"status"` // 状态：1: 上架  2: 下架  3:待上架   default:3
	IsDel      bool      `json:"is_del"`
	SPU        string    `json:"spu"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	PublishAt  time.Time `json:"publish_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
