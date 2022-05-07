/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 19:53
  @Notice:  普通消息
*/

package FriendMsg

import (
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

type Text struct {
	Msg string
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
