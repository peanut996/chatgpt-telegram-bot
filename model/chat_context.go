package model

type ChatContext struct {
	Question string
	UserID   string
	Model    string
}

func NewChatContext(question string, userID string, model string) ChatContext {
	return ChatContext{
		Question: question,
		UserID:   userID,
		Model:    model,
	}
}
