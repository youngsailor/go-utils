/**
 * @Author: 591245853@qq.com
 * @Description:
 * @File: small_tool
 * @Date: 2022/10/12 15:22
 */

package utils

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

// InArray 判断字符串在数组中
func InArray[T comparable](value T, arr []T) bool {
	// 初始化map
	var set = make(map[T]struct{})
	// 将list内容传递进map,只根据key判断，所以不需要关心value的值，用struct{}{}表示
	for _, value := range arr {
		set[value] = struct{}{}
	}
	if _, ok := set[value]; ok {
		return true
	} else {
		return false
	}
}

// CutStr 截取指定长度字符串
func CutStr(str string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(str); i++ {
		_, size = utf8.DecodeRuneInString(str[n:])
		n += size
	}
	return str[:n]
}

// GetMapKeys 获取map的key
func GetMapKeys[T comparable](m map[T]any) []T {
	j := 0
	keys := make([]T, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

// GetRandomBoth 获取随机数  数字和文字
func GetRandomBoth(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetDayStartTime 获取当天零点时间
func GetDayStartTime() int64 {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC).Unix()
}

// GetDayEndTime 获取当天最后时间
func GetDayEndTime() int64 {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, time.UTC).Unix()
}

// FilterEmoji 过滤emoji表情
func FilterEmoji(content string) string {
	newContent := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		}
	}
	if newContent == "" {
		newContent = "-"
	}
	return newContent
}

// SecondsTimer 可以获取剩余时间的timer
type SecondsTimer struct {
	timer *time.Timer
	end   time.Time
}

func NewSecondsTimer(t time.Duration) *SecondsTimer {
	return &SecondsTimer{time.NewTimer(t), time.Now().Add(t)}
}

func (s *SecondsTimer) Reset(t time.Duration) {
	s.timer.Reset(t)
	s.end = time.Now().Add(t)
}

func (s *SecondsTimer) Stop() {
	s.timer.Stop()
}

func (s *SecondsTimer) TimeRemaining() time.Duration {
	return s.end.Sub(time.Now())
}
