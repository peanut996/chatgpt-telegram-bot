package tip

var (
	BotPingTip    = "pong"
	UnknownCmdTip = "Unknown command, please send /start to start a chat \n\n" +
		"🔥未知命令，请发送 /start 来开始聊天"
	BotStartTip = "Please reply to this message with your question, and I will respond to you. \n" +
		"You can also private message me or add me to your group chat. \n" +
		"If you want to chat with me, use the command '/gpt' followed by a space and your question, for example, '/gpt How is the weather today?'. For conversations with GPT-4, please use the command '/gpt4'.\n\n" +
		"😊请在这条消息下回复你的问题，我会回复你的 \n" +
		"🔥你也可以私聊我或者把我加到你的群组聊天 \n" +
		"🤖如果你想和我聊些什么「 /gpt 」+ 空格 + 你的问题，如【/gpt 今天天气怎么样?】. GPT-4对话请使用「 /gpt4 」"

	GPTLackTextTipTemplate = "`/%s` + blank + your question.\n\n" +
		"😊「 /%s 」+ 空格 + 你的问题"

	QueueTipTemplate = "Current queue count: %d, please wait wait patiently\n\n" +
		"🔥当前排队人数：%d, 请耐心等候\n\n"
)
