package middleware

import "net/http"

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		if r.Header.Get("token") == "123456" {
			next(w, r)
			return
		}
		w.Write([]byte("Unauthorized")) // Return unauthorized response

		// Passthrough to next handler if need
	}
}
