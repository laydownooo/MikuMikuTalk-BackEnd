package logic

import (
	"context"
	"errors"
	"fmt"

	"im_server/im_auth/auth_api/internal/svc"
	"im_server/im_auth/auth_api/internal/types"
	"im_server/im_auth/whitelist"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 认证接口
func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp string, err error) {
	if whitelist.IsInList(req.ValidPath, l.svcCtx.Config.WhiteList) {
		return "ok", nil
	}
	if req.Authorization == "" {
		err = errors.New("认证失败")
		return
	}
	payload, err := jwts.ParseToken(req.Authorization, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}
	//从Redis中找一下，能不能找到，找到说明注销了，找不到说明没注销
	_, err = l.svcCtx.RDB.Get(fmt.Sprintf("logout_%s", payload.Nickname)).Result()
	//如果找到了相关数据，那就注销了，直接返回认证失败
	if err == nil {
		err = errors.New("认证失败")
		return
	}
	return "ok", nil

}
