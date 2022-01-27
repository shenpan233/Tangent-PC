/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/26 13:29
  @Notice:  消息回执
*/

package GroupMsg

import (
	"Tangent-PC/utils/GuBuffer"
)

const (
	Receipt = 0x29
)

//ReadMsg 消息已读
func ReadMsg(GroupCode uint64, MsgSeq uint32) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(Receipt) //事件类型
		pack.SetUint32(uint32(GroupCode))
		pack.SetUint8(0x02)
		pack.SetUint32(MsgSeq)
	})
}

//UnReadMsg 处理消息已读
func UnReadMsg(bin []byte) bool {
	pack := GuBuffer.NewGuUnPacket(bin)
	return pack.GetUint8() == 0
}
