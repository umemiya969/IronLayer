package core

var banned = make(map[string]bool)

func IsBanned(ip string) bool {
	return banned[ip]
}
