package httpserver

import (
	"github.com/nathanwdavis/histri"
	"http"
)

type HttpCaptureHandler struct {
}

func (h *HttpCaptureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
