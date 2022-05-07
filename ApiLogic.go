/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:48
  @Notice:  内部业务接口
*/

package Tangent_PC

import (
	"fmt"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Http/HttpGroup/QunInfo"
	GroupMsg "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Receive"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	penguin_http "github.com/shenpan233/penguin-http"
	"strconv"
	"sync"
	"time"
)

//getGroupsInfo 获取本号码群列表以及成员信息 建议写入持久缓存
func (this *TangentPC) getGroupsInfo() {
	qunKey := this.info.SelfWebKey.WebSiteKeys[model.WebQun]
	GroupList := QunInfo.GetGroupList(this.info.Account, this.info.SelfWebKey.Skey, qunKey)
	this.cache.groupList = GroupList
	lock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for Group, _ := range GroupList {
		wg.Add(1)
		go func(Group uint64) {
			member := QunInfo.GetGroupMembers(this.info.Account, strconv.Itoa(int(Group)), this.info.SelfWebKey.Skey, qunKey)
			lock.Lock()
			this.cache.member[Group] = member //使用json方式持久储存
			lock.Unlock()
			wg.Done()
			//GuLog.Notice(this.cache.member[Group])
		}(Group)
		time.Sleep(50 * time.Millisecond) //防止频繁
	}
	wg.Wait()
	//GuLog.Warm(this.cache.member)
	GuLog.NoticeF("[QQ=%d]Finish Loading GroupNum => %d", this.info.LongUin, len(this.cache.groupList))
}

//GetServerMsg 读取系统信息
func (this *TangentPC) GetServerMsg(Cmd, subCmd uint16, seq uint16, MsgInfo, data []byte) {
	go func() {
		buffer := this.packetCommonNoSeq(Cmd, seq, MsgInfo)
		this.udper.Send(&buffer)
	}()
	switch subCmd {
	case 0x00_52:
		Msg := GroupMsg.GroupMsg(data)
		Msg.Account = this.info.LongUin
		Msg.GroupName = this.GetJoinedGroupName(Msg.GroupUin)
		Msg.FromName = this.GetGroupMemberCardFromCache(Msg.GroupUin, Msg.SenderUin)
		if this.hook.GroupMsg != nil {
			go this.ReadGroupMsg(Msg.GroupUin, Msg.MsgSeq)
			go this.hook.GroupMsg(Msg)
		}
		break

	}
}

//refreshClient 刷新ClientKey
func (this *TangentPC) refreshClient() {
	ssoSeq, buffer := this.pack001D(0x11)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		//GuLog.Error("refreshClient", util.BinToHex(bin))
	}
}

//refreshHttpConnSig 刷新HttpConnSig
func (this *TangentPC) refreshHttpConnSig() bool {
	ssoSeq, buffer := this.pack001D(subCmd0x001DHttpConn)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		if gen := this.unpack001D(bin); gen != nil {
			gen := gen.(map[int8][]byte)
			this.teaKey.HttpConn = gen[0]
			this.sig.BufSigHttpConnToken = gen[1]
			return true
		} else {
			return false
		}
	}
	return false
}

//refreshWebKey 刷新WebKey
func (this *TangentPC) refreshWebKey() {
	ssoSeq, buffer := this.pack001D(subCmd0x001DWebKey)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		ret := this.unpack001D(bin)
		if ret != nil {
			this.info.SelfWebKey = new(model.WebKey)
			*this.info.SelfWebKey = *(ret.(*model.WebKey))
		}
	}

}

//GenHttpConn
func (this *TangentPC) genHttpConn() {
	ssoSeq, buffer := this.pack01BB(subCmd0x01BBHttpConn)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
	}
}

func (this *TangentPC) NewLevelSpeedUp() {
	http := penguin_http.Builder().BaseUrl("https://ti.qq.com").Build()
	result, err := http.GET().
		SetCookieFromMap(map[string]string{
			"uin":   "o0" + this.info.Account,
			"skey":  this.info.SelfWebKey.Skey,
			"p_uin": "o0" + this.info.Account,
			//"p_skey": this.info.SelfWebKey.WebSiteKeys[model.WebT],
			"p_skey": this.info.SelfWebKey.PSkey,
		}).
		Sync("/qqlevel/index")
	if err == nil {
		fmt.Println(result.String())
	} else {
		GuLog.Error(err.Error())
	}
}
