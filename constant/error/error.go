package error

var (
	ChatGPTError = "ChatGPT return error, try later again \n\n" +
		"😇出错了, 稍后重试下吧"
	ChatGPTErrorTemplate = "ChatGPT return error, try later again \n\n" +
		"😇出错了, 稍后重试下吧 \n\n%s"
	ChatGPTEngineNotOnline = "Chatgpt engine is not ready, please wait a moment. \n\n" +
		"😇ChatGPT 引擎还没有准备好，请稍等一下"

	NetworkError = "Network error, please try again later \n\n" +
		"😐网络错误，请稍后再试"

	InternalError = "Internal error, please try again later \n\n" +
		"😐内部错误，请稍后再试"
)
