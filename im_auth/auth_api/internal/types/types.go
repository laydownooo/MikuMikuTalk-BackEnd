// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type OpenLoginInfo struct {
	Name string `json:"name"` //用户名
	Icon string `json:"icon"` //用户头像
	Href string `json:"href"` // 跳转地址
}

type OpenLoginRequest struct {
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserName string `json:"username"`
}
