package util

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// StrFirstToUpper 首字母转换成大写
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}

	strArray := []rune(str)

	if strArray[0] >= 97 && strArray[0] <= 122 {
		strArray[0] -= 32
	}

	return string(strArray)
}

// StrFirstToLower 首字母转换成小写
func StrFirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}

	strArray := []rune(str)

	if strArray[0] >= 65 && strArray[0] <= 90 {
		strArray[0] += 32
	}

	return string(strArray)
}

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}
func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

// Camel2Case
// @Description:
// @param name 驼峰式写法转为下划线写法
// @return string
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// StrCamel 将下划线连接的字串改为驼峰
func StrCamel(str string) string {
	tmp := strings.Split(str, "_")
	var result string
	for _, v := range tmp {
		result += StrFirstToUpper(v)
	}

	return result
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
