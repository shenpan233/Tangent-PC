/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 21:28
  @Notice:  notice
*/

package test

import (
	"fmt"
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	"github.com/shenpan233/Tangent-PC/protocal/Msg/Group"
	"github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Send"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuStr"
	"reflect"
	"regexp"
	"strconv"
	"testing"
)

func BenchmarkMsgBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Send.GroupMsg(256783314, "[At=232233][At=All]大丘丘病了，二丘丘哭~ [Pic={35D6BF1F-EADA-FAA3-EF71-827AC0E7A542}.jpg]")
	}
}

func TestMsgLen(t *testing.T) {
	regularAt := fmt.Sprintf(`\%s[0-9]{5,12}\%s|\%sAll\%s`, Msg.FormatAt, Msg.FormatEnd, Msg.FormatAt, Msg.FormatEnd)
	regularPic := fmt.Sprintf(`\%s\{[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\}\.[jpginf]{3}%s`, Msg.FormatPic, Msg.FormatEnd)
	CommonText := `.[^[{\\]{0,500}`
	must := regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, regularPic, regularAt, CommonText))
	AllFound := must.FindAllString("[At=232233][At=All]大丘丘病了，二丘丘哭~ [Pic={35D6BF1F-EADA-FAA3-EF71-827AC0E7A542}.jpg]", -1)
	ret := make([]interface{}, 0)
	for _, SubStr := range AllFound {
		//fmt.Print(SubStr)
		if isAt, _ := regexp.MatchString(regularAt, SubStr); isAt {
			uin, _ := strconv.Atoi(GuStr.Between(SubStr, Msg.FormatAt, Msg.FormatEnd))
			ret = append(ret, &Group.Common{
				IsAt:  true,
				AtUin: uint32(uin),
			})
		} else if isPic, _ := regexp.MatchString(regularPic, SubStr); isPic {
			ret = append(ret, &Group.Pic{
				Guid: GuStr.Between(SubStr, Msg.FormatPic, Msg.FormatEnd),
			})
		} else {
			ret = append(ret, &Group.Common{
				IsAt: false,
				Msg:  SubStr,
			})
		}
	}
	//

	for _, subCall := range ret {
		fmt.Println(subCall)
		fmt.Println(util.BinToHex(reflect.ValueOf(subCall).MethodByName("Marshal").Call(nil)[0].Bytes()))
	}

}
