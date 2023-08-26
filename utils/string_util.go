package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func SplitMessageByMaxSize(msg string, maxSize int) []string {
	var msgs []string
	currentMsg := msg

	if len(currentMsg) <= maxSize {
		msgs = append(msgs, currentMsg)
		return msgs
	}

	for len(currentMsg) > maxSize {
		msgs = append(msgs, currentMsg[:maxSize])
		currentMsg = currentMsg[maxSize:]
	}
	msgs = append(msgs, currentMsg)
	return msgs
}

func GenerateInvitationCode(size int) (string, error) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, size)
	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		result[i] = chars[index.Int64()]
	}
	return string(result), nil
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
