package commonutils

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

const (
	md5Pattern = `^[a-fA-F0-9]{32}$`
)

// IsMD5Format 判断输入的字符串是不是md5格式的字符串
func IsMD5Format(paramValue string) bool {
	pattern := md5Pattern //`^[a-fA-F0-9]{32}$`
	match, _ := regexp.MatchString(pattern, paramValue)
	return match
}

// EasyMD5 生成MD5字符串
func EasyMD5(paramArgs ...string) string {
	data := []byte(StringCat(paramArgs...))
	return fmt.Sprintf("%x", md5.Sum(data))
}
