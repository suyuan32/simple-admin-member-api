package middleware

import (
	"net/http"
	"strings"
)

type VIPMiddleware struct {
}

func NewVIPMiddleware() *VIPMiddleware {
	return &VIPMiddleware{}
}

func (m *VIPMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the role id
		roleIds := r.Context().Value("rankId").(string)

		if !strings.Contains(roleIds, "002") {

		}

		next(w, r)
	}
}
