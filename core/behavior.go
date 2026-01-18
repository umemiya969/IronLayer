package core

import "strings"

func BehaviorCheck(ctx *RequestContext, cfg *Config) {
	for _, p := range cfg.Rules.BlockPath {
		if strings.Contains(ctx.Path, p) {
			ctx.Score += 10
		}
	}
}
