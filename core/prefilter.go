package core

import (
	"net"
	"strings"
)

func PreFilter(ctx *RequestContext, cfg *Config) {

	// block empty IP
	if ctx.IP == "" {
		ctx.Score += 20
		return
	}

	// block localhost
	if ctx.IP == "127.0.0.1" || ctx.IP == "::1" {
		ctx.Score += 20
		return
	}

	// block private network
	ip := net.ParseIP(ctx.IP)
	if ip != nil {
		if ip.IsPrivate() {
			ctx.Score += 20
			return
		}
	}

	// suspicious UA
	if strings.TrimSpace(ctx.UA) == "" {
		ctx.Score += 3
		ctx.reason = "prefilter"
	}
}
