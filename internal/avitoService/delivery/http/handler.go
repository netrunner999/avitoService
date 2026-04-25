package http

import (
	"net/http"
	avitoService "avitoService/internal/avitoService"
)

// структура http-слоя, хранит ссылку на бизнес логику
type handler struct {
	u avitoService.UseCase
}

func New(u avitoService.UseCase) *handler {
	return &handler{u: u}
}

// хэндлер Get /test
func (h *handler) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg, err := h.u.GetTestMessage(r.Context())
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	}
}
