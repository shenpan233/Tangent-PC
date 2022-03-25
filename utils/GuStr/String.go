/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 23:04
  @Notice:  字符串的操作
*/

package GuStr

import (
	"regexp"
	"strings"
)

var (
	mustPassage = regexp.MustCompile(".")
)

func Between(str, Start, End string) string {
	s := strings.Index(str, Start)
	if s < 0 {
		return ""
	}
	s += len(Start)
	e := strings.Index(str[s:], End)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

//GetPassLen 字数统计 英文1:1 中文1:1
func GetPassLen(passage string) int {
	must := mustPassage
	m := must.FindAllStringSubmatchIndex(passage, -1)
	return len(m)
}
