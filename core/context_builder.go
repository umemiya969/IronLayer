package core

import (
	"net"
	"net/http"
	"strings"
)

func BuildContext(r *http.Request) *RequestContext {
	ip := r.RemoteAddr
	if strings.Contains(ip, ":") {
		ip, _, _ = net.SplitHostPort(ip)
	}

	return &RequestContext{
		IP:      ip,
		Path:    r.URL.Path,
		Method:  r.Method,
		UA:      r.UserAgent(),
		Request: r,
	}
}
