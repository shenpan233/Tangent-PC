/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/19 14:38
  @Notice:  获取群列表
*/

package QunInfo

import (
	json "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"github.com/shenpan233/Tangent-PC/protocal/Http"
)

type groupListJson struct {
	Ec      int    `json:"ec"`
	Errcode int    `json:"errcode"`
	Em      string `json:"em"`
	Join    []struct {
		Gc    int    `json:"gc"`
		Gn    string `json:"gn"`
		Owner int    `json:"owner"`
	} `json:"join"`
	Create []struct {
		Gc    int    `json:"gc"`
		Gn    string `json:"gn"`
		Owner int    `json:"owner"`
	} `json:"create"`
}
type GroupList struct {
	GroupUin uint64
	Name     string
}

//GetGroupList 获取加入和创建的群列表[不进行区分]
func GetGroupList(uin string, skey, baseKey string) (GroupLists map[uint64]string) {
	req := gorequest.New().Post("https://qun.qq.com/cgi-bin/qun_mgr/get_group_list")
	req.Send("bkn="+Http.GenGtk(skey)).
		Set("User-Agent", Http.UserAgent).
		AddCookies(Http.PkgCommonCookies(uin, skey, baseKey)).
		EndBytes(func(response gorequest.Response, body []byte, errs []error) {
			GroupListRoot := new(groupListJson)
			GroupLists = make(map[uint64]string)
			if err := json.Unmarshal(body, GroupListRoot); err == nil {
				for _, joinGroup := range GroupListRoot.Join {
					GroupLists[uint64(joinGroup.Gc)] = joinGroup.Gn
				}
				for _, createGroup := range GroupListRoot.Create {
					GroupLists[uint64(createGroup.Gc)] = createGroup.Gn
				}
			}
		})

	return
}
