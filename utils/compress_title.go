package utils

import "regexp"

//压缩标题，去掉空格
func CompressTitle(str string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}
