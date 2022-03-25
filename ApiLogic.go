/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:48
  @Notice:  内部业务
*/

package Tangent_PC

import (
	"bytes"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Http/HttpGroup/QunInfo"
	GroupMsg "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Receive"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"runtime"
	"strconv"
	"sync"
	"time"
)

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

//finishLogin 登录完成后的操作
func (this *TangentPC) finishLogin() {
	//刷新token
	this.refreshClient()
	this.refreshHttpConnSig()
	this.refreshWebKey()
	this.getGroupsInfo()
	//绑定接收器
	this.udper.UdpRecv = this.Recv
	this.handle = map[uint16]unpack{
		0x00_17: this.unpack0017,
	}
}

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
			GuLog.Notice(this.cache.member[Group])
		}(Group)
		time.Sleep(50 * time.Millisecond) //防止频繁
	}
	wg.Wait()
	GuLog.Warm(this.cache.member)
	GuLog.NoticeF("[QQ=%d]Finish Loading GroupNum => %d", this.info.LongUin, len(this.cache.groupList))
}

//GetServerMsg 读取系统信息
func (this *TangentPC) GetServerMsg(Cmd uint16, seq uint16, MsgInfo, data []byte) {
	go func() {
		buffer := this.pack0017(seq, MsgInfo)
		this.udper.Send(&buffer)
	}()
	switch Cmd {
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

//Recv 数据包接收
func (this *TangentPC) Recv(Cmd uint16, seq uint16, pack *GuBuffer.GuUnPacket) {
	//异常捕获
	defer func() {
		if err := recover(); err != nil {
			lineW := bytes.NewBufferString("")
			for i := 0; i < 2; i++ {
				_, file, line, _ := runtime.Caller(3 + i)
				lineW.WriteString(file + ":" + strconv.Itoa(line))
				lineW.WriteString("\n")
			}
			GuLog.ErrorF("%s\n%v\n", lineW.String(), err)
		}
	}()
	pack.GetBin(3)
	pack = GuBuffer.NewGuUnPacket(util.Decrypt(this.teaKey.SessionKey, pack.GetAll()))
	if event := this.handle[Cmd]; event != nil {
		event(seq, pack.GetAll())
	} else {
		//GuLog.Info("ReCv", "QQ:[%d],Cmd=0x%X,Buff=%X", this.info.LongUin, Cmd, pack.GetAll())
	}

}
