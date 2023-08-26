package cmd

var (
	START = "start"
	PING  = "ping"
	GPT   = "gpt"

	DOWN = "downgrade"

	PPROF = "pprof"

	INVITE = "invite"

	COUNT = "count"

	QUERY = "query"

	GPT4 = "gpt4"

	DONATE = "donate"

	VIP = "vip"

	PUSH = "push"

	STATUS = "status"

	ACCESS = "access"

	_ = "cmd"
)

func IsBotCmd(cmd string) bool {
	switch cmd {
	case START, PING, GPT, DOWN,
		PPROF, INVITE, COUNT, QUERY,
		GPT4, DONATE, PUSH, STATUS, VIP,
		ACCESS:
		return true
	default:
		return false
	}
}
