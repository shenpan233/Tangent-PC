/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/24 12:06
  @Notice:  聊天相关的Api
*/

package Tangent_PC

import (
	"errors"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	GroupMsg "github.com/shenpan233/Tangent-PC/protocal/Msg/Group/Receive"
	"github.com/shenpan233/Tangent-PC/protocal/Protobuf/im/cmd0x0002"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

//RevokeGroupMessage 撤回消息	(:要有管理员权限
func (this *TangentPC) RevokeGroupMessage(GroupCode uint64, MsgSeq, MsgID uint32) error {
	ssoSeq, buffer := this.pack0x3f7(GroupCode, MsgSeq, MsgID)
	bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer)
	if bin != nil {
		return this.unpack0x3f7(bin)
	} else {
		return errors.New("revoked GroupMessage fail,No bytes was returned")
	}
}

//ReadGroupMsg 置群消息已读
//	内部会自动调用不用管
func (this *TangentPC) ReadGroupMsg(GroupCode uint64, MsgSeq uint32) bool {
	ssoSeq, buffer := this.pack0002(GroupMsg.ReadMsg(GroupCode, MsgSeq))
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		isSuc, _ := this.unpack0002(bin)
		return isSuc
	}
	return false
}

//SendGroupMsg 发送群消息
//	GroupCode 群号
//	Msg 	  消息内容
func (this *TangentPC) SendGroupMsg(GroupCode uint64, Msg ...Msg.Builder) (Code bool, MsgSeq uint32) {
	ssoSeq, buffer := this.pack0002(GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(_0x0002Send) //事件类型
		pack.SetUint32(uint32(GroupCode))
		pack.SetToken(GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetUint16(1)
			pack.SetUint8(uint8(1))
			pack.SetUint8(uint8(0))
			pack.SetBytes([]byte{0x00, 0x00})                                     //RandomSeq
			pack.SetBytes([]byte{0x00, 0x00, 0x00, 0x00})                         //固定空白4 字节
			pack.SetBytes([]byte{0x4D, 0x53, 0x47, 0x00, 0x00, 0x00, 0x00, 0x00}) //MSG
			pack.SetUint32(uint32(util.GetServerCurTime()))
			pack.SetUint32(util.GetRand32())
			pack.SetBytes([]byte{0x00}) //一个分隔符
			//TODO GroupMsg:Font自定义
			font := model.Font{
				Red:      0,
				Blue:     0,
				Green:    0,
				Size:     0x0A, //老年人都看得见
				Encoding: model.FontEncodingUTF8,
				FontName: model.FontNameMicrosoftYaHei,
			}
			pack.SetBytes(font.ToBytes())
			pack.SetBytes([]byte{0x00, 0x00})
			//构造消息
			for _, builder := range Msg {
				pack.SetBytes(builder.Generate())
			}
		}))

	}),
	)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
		isSuc, Recall := this.unpack0002(bin)
		if isSuc {
			msg := Recall.(cmd0x0002.SendGroupMsg)
			return true, msg.GetMsgSeq()
		}
	}
	return false, 0
}

func (this *TangentPC) SendFriendMsg(FriendUin uint64, Msg ...Msg.Builder) (Code bool, MsgSeq uint32) {
	if len(Msg) == 0 {
		return false, 0
	}
	ssoSeq, buffer := this.pack00CD(FriendUin, Msg...)
	if bin := this.udper.SendAndGet(ssoSeq, WaitTime, &buffer); bin != nil {
	}
	return true, 0
}

//GetJoinedGroupName 获取已加入的群列表
func (this *TangentPC) GetJoinedGroupName(GroupUin uint64) string {
	return this.cache.groupList[GroupUin]
}

//GetGroupMemberCardFromCache 从缓存获取群员的群名片
func (this *TangentPC) GetGroupMemberCardFromCache(GroupUin, uin uint64) string {
	if member := this.cache.member[GroupUin][uin]; member != nil {
		return member.Name
	}
	return ""
}
