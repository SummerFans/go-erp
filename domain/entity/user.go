package entity

type User struct {
}

type UserAuth struct {
	ID           int    // 授权ID
	UserID       int    // 用户ID
	IdentityType string // 登录类型
	Identity     string // 标志
	Credential   string // 密码凭证
}

// 用户级别信息
type UserInfo struct {
	ID       int    // 用户信息ID
	UserID   int    // 用户ID
	UserName string // 用户名称
	Gender   int    // 0: 保密  1: 男  2: 密码
}

// 登录日志记录
type UserRecord struct {
}
