package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hello01/internal/logic"
	"hello01/internal/svc"
	"hello01/internal/types"
)

func Hello01Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 请求参数
		var req types.Request
		// 解析参数
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 注入serviceContext
		l := logic.NewHello01Logic(r.Context(), svcCtx)
		// 业务逻辑实现
		resp, err := l.Hello01(&req)
		// 返回响应结果
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
