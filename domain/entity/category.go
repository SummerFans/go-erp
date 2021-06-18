package entity

import "time"

type Category struct {
	CategoryID   int        `json:"category_id"`                      // 分类ID
	CategoryName string     `json:"category_name" binding:"required"` // TODO 分类名称 (根据语言获取不同名称)
	ParentID     int        `json:"parent_id"`                        // 父级ID
	CategoryCode string     `json:"category_code"`                    // 分类Code
	CreatedAt    *time.Time `json:"created_at"`                       // 创建时间
	DeletedAt    *time.Time `json:"deleted_at"`                       // 删除时间
	UpdatedAt    *time.Time `json:"updated_at"`                       // 更新时间
}

type CategoryLang struct {
	CategoryID int `json:"category_id"`
}

// 获取分类类表
type CategoryGetAllParams struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
