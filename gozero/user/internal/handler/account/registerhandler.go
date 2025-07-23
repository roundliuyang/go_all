package account

import (
	"net/http"
	"user/biz"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/internal/logic/account"
	"user/internal/svc"
	"user/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	// 你可以将统一的错误处理添加到 RegisterHandler 函数中，使用 httpx.SetErrorHandler 进行全局错误处理。
	// 只要你调用 httpx.ErrorCtx 进行错误响应时，SetErrorHandler 就会自动生效。
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *biz.Error:
			// 当错误为自定义的 biz.Error 类型时，返回 200 状态码
			return http.StatusOK, biz.Fail(e)
		default:
			// 对于其他错误，返回 500 状态码
			return http.StatusInternalServerError, nil
		}
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 调用业务逻辑层
		l := account.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			// 错误处理
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 成功响应
			httpx.OkJsonCtx(r.Context(), w, biz.Success(resp))
		}
	}
}
