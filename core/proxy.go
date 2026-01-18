package core

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewIronLayer(cfg *Config) http.Handler {
	target, _ := url.Parse(cfg.Server.Backend)

	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := BuildContext(r)

		PreFilter(ctx, cfg)
		RateLimit(ctx, cfg)
		BehaviorCheck(ctx, cfg)
		Decision(ctx)

		if ctx.Decision == "block" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		proxy.ServeHTTP(w, r)
	})
}
