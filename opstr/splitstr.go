package opstr

import (
	"regexp"
	"strings"
)

// SplitBySpace split string and get one string
// 获取多个空格分隔的第几个字符，th从0开始计算
func SplitBySpace(str string, th int) string {
	str = strings.TrimSpace(str)
	reg := regexp.MustCompile("\\s{2,}")
	str = reg.ReplaceAllString(str, " ")

	strs := strings.Split(str, " ")
	if len(strs) > th || th >= 0 {
		return strs[th]
	}

	return ""
}
