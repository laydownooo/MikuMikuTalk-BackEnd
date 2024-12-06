// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"im_server/im_chat/chat_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 聊天记录接口
				Method:  http.MethodGet,
				Path:    "/history",
				Handler: chatHistoryHandler(serverCtx),
			},
			{
				// 最近聊天会话列表
				Method:  http.MethodGet,
				Path:    "/session",
				Handler: chatSessionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/chat"),
	)
}
