package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	guestPrefix = "访客"
	charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// 生成访客ID，格式为"访客"+5位随机字母数字
func GenerateGuestID() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return guestPrefix + string(b)
}

// 检查是否是访客ID
func IsGuestID(id string) bool {
	return strings.HasPrefix(id, guestPrefix)
}

// 标准化用户名，处理冲突
func NormalizeUsername(username string, isConflict func(string) bool) string {
	if username == "" {
		return ""
	}

	if !isConflict(username) {
		return username
	}

	// 添加随机后缀解决冲突，使用5位随机字符保持统一
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return username + string(b)
}
