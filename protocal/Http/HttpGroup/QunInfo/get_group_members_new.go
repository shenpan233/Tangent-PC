/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/19 11:07
  @Notice:  获取群成员信息
*/

package QunInfo

import (
	json "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Http"
	"strconv"
)

type getGroupMembersNew struct {
	Ec      int    `json:"ec"`
	Errcode int    `json:"errcode"`
	Em      string `json:"em"`
	ExtNum  int    `json:"ext_num"`
	Level   int    `json:"level"`
	GMemNum int    `json:"gMemNum"`
	Owner   int    `json:"owner"`
	Type    int    `json:"type"`
	Mems    []mem
	Cards   interface{} `json:"cards"`
}

type mem struct {
	U uint64 `json:"u"`
	G int    `json:"g"`
	N string `json:"n"`
	B int    `json:"b"`
}

//GetGroupMembers 获取群成员,请少用
func GetGroupMembers(uin, group string, skey, baseKey string) (member map[uint64]*model.GroupMember) {
	req := gorequest.New().Post("https://qinfo.clt.qq.com/cgi-bin/qun_info/get_group_members_new")
	req.AddCookies(Http.PkgCommonCookies(uin, skey, baseKey)).
		Send(map[string]string{
			"gc":  group,
			"bkn": Http.GenGtk(skey),
		}).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("User-Agent", Http.UserAgent).
		EndBytes(func(_ gorequest.Response, body []byte, _ []error) {
			var getGroupMembersNew getGroupMembersNew
			member = make(map[uint64]*model.GroupMember)
			if err := json.Unmarshal(body, &getGroupMembersNew); err == nil {
				//无错误
				if getGroupMembersNew.Cards != nil {
					cards := getGroupMembersNew.Cards.(map[string]interface{})
					for _, m := range getGroupMembersNew.Mems {
						tmpCard := cards[strconv.Itoa(int(m.U))]
						card := ""
						if tmpCard != nil {
							card = tmpCard.(string)
						}
						member[m.U] = &model.GroupMember{
							Name: m.N,
							Card: card,
						}
					}
				} else {
					for _, m := range getGroupMembersNew.Mems {
						member[m.U] = &model.GroupMember{
							Name: m.N,
						}
					}
				}
			}
			return
		})
	return member
}
