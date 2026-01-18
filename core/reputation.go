package core

var reputation = make(map[string]int)

func UpdateReputation(ctx *RequestContext) {
	reputation[ctx.IP] += ctx.Score
}
