/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 19:53
  @Notice:  普通消息
*/

package FriendMsg

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	"github.com/shenpan233/Tangent-PC/protocal/Protobuf/im/cmd0x0002"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuStr"
	"strconv"
)

type Text struct {
	Msg string `json:"msg"`
}

func (t *Text) Generate() []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetLitTlv(Msg.TypeText, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetLitTlv(Msg.CommonMsg, []byte(t.Msg))
		}))
	})
}

func (t *Text) Decode() string {
	return t.Msg
}

type At struct {
	Uin  uint64 `json:"account"`
	Card string
}

func (a *At) Generate() []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetLitTlv(Msg.TypeText, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			var tips string
			if a.Uin == 0 {
				tips = "@全体成员"
			} else {
				tips = "@" + a.Card
			}
			pack.SetLitTlvS(Msg.CommonMsg, tips)
			pack.SetLitTlv(Msg.CommonAt, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
				pack.SetBytes([]byte{0x00, 0x01})
				pack.SetUint32(uint32(GuStr.GetPassLen(tips)))
				pack.SetBytes([]byte{0x00})
				pack.SetUint32(uint32(a.Uin))
				pack.SetBytes([]byte{0x00, 0x00})
			}))
		}))
	})
}

func (a *At) Decode() string {
	return Msg.FormatAt + strconv.Itoa(int(a.Uin)) + Msg.FormatEnd
}

type Reply struct {
	Uin       uint64 `json:"uin"`
	Card      string `json:"card"`
	MsgSeq    uint32 `json:"msgSeq"`
	SendTime  uint64 `json:"sendTime"`
	GroupCode uint64 `json:"groupCode"`
}

func (this *Reply) Generate() []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes((&At{
			Uin:  this.Uin,
			Card: this.Card,
		}).Generate()) //At对方
		pack.SetBytes((&Text{
			Msg: " ",
		}).Generate()) //空格
		pack.SetLitTlv(Msg.TypeReply, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			reply := cmd0x0002.Reply{
				Targer: &cmd0x0002.ReplyReplyTarger{
					MsgSeq:   &this.MsgSeq,
					FromUin:  &this.Uin,
					SendTime: &this.SendTime,
					Tag4:     proto.Uint32(cmd0x0002.Default_ReplyReplyTarger_Tag4),
					ShowMsg: &cmd0x0002.ShowMsg{
						Show: &cmd0x0002.ShowMsg_Showed{
							Text: proto.String("回复消息"),
						},
					},
					Tag6:      proto.Uint64(cmd0x0002.Default_ReplyReplyTarger_Tag6),
					GroupCode: &this.GroupCode,
				},
			}
			marshal, _ := reply.Marshal()
			pack.SetLitTlv(0x01, marshal)
		}))
	})
}

func (this *Reply) Decode() string {
	return fmt.Sprintf(Msg.FormatReply+Msg.FormatEnd, strconv.Itoa(int(this.Uin)), strconv.Itoa(int(this.MsgSeq)), strconv.Itoa(int(this.SendTime)))
}
