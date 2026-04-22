package dto

// UserLoginRequest 登录实体
type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	ID       uint64 `json:"id"`       // 用户ID
	Username string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 角色
	Token    string `json:"token"`    // token
}

type UserInfoResponse struct {
	ID       uint64 `json:"id"`       // 用户ID
	Username string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 角色
}
