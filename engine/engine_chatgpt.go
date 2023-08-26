package engine

import (
	"chatgpt-telegram-bot/cfg"
	"chatgpt-telegram-bot/constant/config"
	botError "chatgpt-telegram-bot/constant/error"
	"chatgpt-telegram-bot/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ChatGPTEngine struct {
	client *http.Client

	baseUrl string

	alive bool
}

func NewChatGPTEngine() *ChatGPTEngine {
	return &ChatGPTEngine{}
}

func (e *ChatGPTEngine) Init(cfg *cfg.Config) error {
	e.client = &http.Client{}
	e.baseUrl = fmt.Sprintf("http://%s:%d", cfg.EngineConfig.Host, cfg.EngineConfig.Port)

	go e.checkChatGPTEngine()
	return nil
}

func (e *ChatGPTEngine) Alive() bool {
	resp, err := http.Get(e.baseUrl + "/ping")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

// Chat is the method to chat with ChatGPT engine
func (e *ChatGPTEngine) chat(ctx model.ChatContext) (string, error) {
	sentence, userID, gptModel := ctx.Question, ctx.UserID, ctx.Model
	log.Println("[ChatGPT] send request to chatgpt, text: ", sentence)

	if !e.Alive() {
		return botError.ChatGPTEngineNotOnline, nil
	}

	encodeSentence := url.QueryEscape(sentence)
	e.client.Timeout = time.Duration(config.ChatGPTTimeoutSeconds) * time.Second
	queryString := fmt.Sprintf("/chat?user_id=%s&sentence=%s&model=%s", userID, encodeSentence, gptModel)
	resp, err := e.client.Get(e.baseUrl + queryString)
	if resp == nil {
		return "", errors.New(botError.NetworkError)
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println("[ChatGPT] chatgpt engine error: ", err)
		return "", errors.New(botError.NetworkError)
	}
	if resp.StatusCode != 200 {
		log.Println("[ChatGPT] chatgpt engine fail, status code: ", resp.StatusCode)
		return "", errors.New(botError.ChatGPTError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ChatGPT] chatgpt engine error: ", err)
		return "", errors.New(botError.InternalError)
	}
	data := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("[ChatGPT] unmarshal chatgpt response error: %s, resp: %s\n",
			err, string(body))
		return "", errors.New(botError.InternalError)
	}

	if msg, ok := data["message"].(string); ok && msg != "" {
		return msg, nil
	}
	if detail, ok := data["detail"].(string); ok && detail != "" {
		return "", fmt.Errorf(botError.ChatGPTErrorTemplate, detail)
	}
	return "", errors.New(botError.ChatGPTError)
}

func (e *ChatGPTEngine) Chat(ctx model.ChatContext) (string, error) {
	resp, err := e.chat(ctx)

	isNetworkError := strings.Contains(resp, "SSLError") ||
		strings.Contains(resp, "RemoteDisconnected") ||
		strings.Contains(resp, "ConnectionResetError")
	if isNetworkError {
		return "", errors.New(botError.NetworkError)
	}

	if err == nil && resp != "" {
		return resp, nil
	}

	if err != nil {
		return "", err
	}
	return botError.ChatGPTError, err
}

func (e *ChatGPTEngine) checkChatGPTEngine() {
	for {
		status := e.Alive()
		if !status {
			log.Println("[HealthCheck] chatgpt engine is not ready")
			e.alive = false
		} else {
			e.alive = true
		}
		time.Sleep(10 * time.Second)
	}
}
