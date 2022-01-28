/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 21:42
  @Notice:  消息结构化工具
*/

package Send

import (
	"fmt"
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	"github.com/shenpan233/Tangent-PC/protocal/Msg/Group"
	"github.com/shenpan233/Tangent-PC/utils/GuStr"
	"regexp"
	"strconv"
)

//用内存交互性能
var (
	regularAt         = fmt.Sprintf(`\%s[0-9]{5,12}\%s|\[At=All\]`, Msg.FormatAt, Msg.FormatEnd)
	regularPic        = fmt.Sprintf(`\%s\{[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\}\.[jpginf]{3}%s`, Msg.FormatPic, Msg.FormatEnd)
	regularCommonText = `.[^[{\\]{0,300}`
	match             = regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, regularPic, regularAt, regularCommonText))
	matchAt           = regexp.MustCompile(regularAt)
	matchPic          = regexp.MustCompile(regularPic)
)

//BuildMsgStructure 结构化构建
func BuildMsgStructure(data string) []interface{} {
	AllFound := match.FindAllString(data, -1)
	ret := make([]interface{}, 0)
	for _, SubStr := range AllFound {
		if matchAt.MatchString(SubStr) {
			uin, _ := strconv.Atoi(GuStr.Between(SubStr, Msg.FormatAt, Msg.FormatEnd))
			ret = append(ret, &Group.Common{
				IsAt:  true,
				AtUin: uint32(uin),
			})
		} else if matchPic.MatchString(SubStr) {
			ret = append(ret, &Group.Pic{
				Guid: GuStr.Between(SubStr, Msg.FormatPic, Msg.FormatEnd),
			})
		} else {
			ret = append(ret, &Group.Common{
				Msg: SubStr,
			})
		}
	}
	return ret
}
