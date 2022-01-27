/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 21:28
  @Notice:  notice
*/

package test

import (
	"Tangent-PC/protocal/Msg"
	"Tangent-PC/utils/Str"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"testing"
)

func BenchmarkMsgBuild(b *testing.B) {
	At := fmt.Sprintf(`\%s[0-9]{5,12}\%s|\[At=All\]`, Msg.FormatAt, Msg.FormatEnd)
	Pic := fmt.Sprintf(`\%s\{[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\}\.[jpginf]{3}%s`, Msg.FormatPic, Msg.FormatEnd)
	CommonText := `.[^[{\\]{0,300}`
	must := regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, Pic, At, CommonText))
	for i := 0; i < b.N; i++ {
		AllFound := must.FindAllString("[At=232233][At=All]大丘丘病了，二丘丘哭~ [Pic={35D6BF1F-EADA-FAA3-EF71-827AC0E7A542}.jpg]", -1)
		ret := make([]interface{}, 0)
		for _, SubStr := range AllFound {
			//fmt.Print(SubStr)
			if isAt, _ := regexp.MatchString(At, SubStr); isAt {
				//fmt.Println(" (:this is an at")
				uin, _ := strconv.Atoi(Str.Between(SubStr, Msg.FormatAt, Msg.FormatEnd))
				ret = append(ret, &Msg.Common{
					IsAt:  true,
					AtUin: uint32(uin),
				})
			} else if isPic, _ := regexp.MatchString(Pic, SubStr); isPic {
				//fmt.Println(" (:this is a beautiful picture")
			} else {
				ret = append(ret, &Msg.Common{
					IsAt: false,
					Msg:  SubStr,
				})
			}
		}
		//

		for _, subCall := range ret {
			reflect.ValueOf(subCall).MethodByName("Marshal").Call(nil)[0].Bytes()
		}
	}
}

func TestMsgLen(t *testing.T) {
	At := fmt.Sprintf(`\%s[0-9]{5,12}\%s|\[At=All\]`, Msg.FormatAt, Msg.FormatEnd)
	Pic := fmt.Sprintf(`\%s\{[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\}\.[jpginf]{3}%s`, Msg.FormatPic, Msg.FormatEnd)
	CommonText := `.[^[{\\]{0,500}`
	must := regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, Pic, At, CommonText))
	AllFound := must.FindAllString("[At=232233][At=All]大丘丘病了，二丘丘哭~ [Pic={35D6BF1F-EADA-FAA3-EF71-827AC0E7A542}.jpg]", -1)
	ret := make([]interface{}, 0)
	for _, SubStr := range AllFound {
		//fmt.Print(SubStr)
		if isAt, _ := regexp.MatchString(At, SubStr); isAt {
			//fmt.Println(" (:this is an at")
			uin, _ := strconv.Atoi(Str.Between(SubStr, Msg.FormatAt, Msg.FormatEnd))
			ret = append(ret, &Msg.Common{
				IsAt:  true,
				AtUin: uint32(uin),
			})
		} else if isPic, _ := regexp.MatchString(Pic, SubStr); isPic {
			//fmt.Println(" (:this is a beautiful picture")
		} else {
			ret = append(ret, &Msg.Common{
				IsAt: false,
				Msg:  SubStr,
			})
		}
	}
	//

	for _, subCall := range ret {
		fmt.Println(reflect.ValueOf(subCall).MethodByName("Marshal").Call(nil)[0].Bytes())
	}

}
