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
	regularReply      = fmt.Sprintf(`\`+Msg.FormatReply+`\`+Msg.FormatEnd, "[0-9]{5,12}", "[0-9]+", "[0-9]{10,13}")
	match             = regexp.MustCompile(fmt.Sprintf(`%s|%s|%s|%s`, regularPic, regularAt, regularReply, regularCommonText))
	matchAt           = regexp.MustCompile(regularAt)
	matchPic          = regexp.MustCompile(regularPic)
	matchReply        = regexp.MustCompile(regularReply)
)

//BuildMsgStructure 结构化构建
func BuildMsgStructure(data string, GroupCode uint64, atName GetGroupMemberCardFromCache) []interface{} {
	AllFound := match.FindAllString(data, -1)
	ret := make([]interface{}, 0)
	for _, SubStr := range AllFound {
		if matchAt.MatchString(SubStr) {
			//群内@
			uin, _ := strconv.Atoi(GuStr.Between(SubStr, Msg.FormatAt, Msg.FormatEnd))
			ret = append(ret, &Group.Common{
				IsAt:  true,
				AtUin: uint32(uin),
				Msg:   atName(GroupCode, uint64(uin)),
			})
		} else if matchPic.MatchString(SubStr) {
			//发送图片
			ret = append(ret, &Group.Pic{
				Guid: GuStr.Between(SubStr, Msg.FormatPic, Msg.FormatEnd),
			})
		} else if matchReply.MatchString(SubStr) {
			uin, _ := strconv.ParseUint(GuStr.Between(SubStr, "FromUin=", ","), 10, 64)
			SendTime, _ := strconv.ParseUint(GuStr.Between(SubStr, "SendTime=", "]"), 10, 64)
			MsgSeq, _ := strconv.ParseUint(GuStr.Between(SubStr, "MsgSeq=", ","), 10, 64)
			reply := &Group.Reply{
				GroupCode: GroupCode,
				FromUin:   uin,
				SendTime:  SendTime,
				MsgSeq:    uint32(MsgSeq),
			}
			ret = append(ret, reply)
		} else {
			ret = append(ret, &Group.Common{
				Msg: SubStr,
			})
		}
	}
	return ret
}
