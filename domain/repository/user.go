package repository

// GoodsRepository ...
type UserRepository interface {

	// 创建用户
	Create() error
	// 获取获取商品列表
	// Get(id int) (*entity.Goods, error)
	// GetAll() ([]*entity.Goods, error)
}
