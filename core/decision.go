package core

func Decision(ctx *RequestContext) {
	if ctx.Score >= 20 {
		ApplyBan(ctx.IP)
		ctx.Decision = "block"
		return
	}

	ctx.Decision = "allow"
}
