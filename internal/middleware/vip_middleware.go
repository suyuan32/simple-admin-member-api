package middleware

import (
	"github.com/suyuan32/simple-admin-member-api/internal/enum/memberrank"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type VipMiddleware struct {
}

func NewVipMiddleware() *VipMiddleware {
	return &VipMiddleware{}
}

func (m *VipMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the rank id
		rankId := r.Context().Value("rankId").(string)

		if !strings.Contains(rankId, memberrank.VIP) {
			httpx.Error(w, errorx.NewApiUnauthorizedError("only VIP user can access the API"))
			return
		}

		next(w, r)
	}
}
