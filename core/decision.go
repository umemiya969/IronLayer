package core

func Decision(ctx *RequestContext) {
	if ctx.Score >= 20 {
		ctx.Decision = "block"
	} else {
		ctx.Decision = "allow"
	}
}
