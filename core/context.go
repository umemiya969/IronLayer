package core

import "net/http"

type RequestContext struct {
	IP       string
	Path     string
	Method   string
	UA       string
	Score    int
	Decision string
	reason   string
	Request  *http.Request
}
