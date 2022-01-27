/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 19:25
  @Notice:  发送消息
*/

package Send

import (
	"Tangent-PC/model"
	Model "Tangent-PC/protocal/Msg/Group"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

//GroupMsg 发送群消息
//	Msg 消息内容
func GroupMsg(GroupUin uint64, Msg string) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(Model.Send) //事件类型
		pack.SetUint32(uint32(GroupUin))
		//part1
		pack.SetToken(GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetUint16(1)
			num := 1 //TODO GroupMsg:分片总数暂时不知道干嘛的
			pack.SetUint8(uint8(num))
			for i := 0; i < num; i++ {
				pack.SetUint8(uint8(i))
				pack.SetBytes(util.GetRandomBin(2)) //RandomSeq
			}
			pack.SetBytes([]byte{0, 0, 0, 0})                                     //固定空白4字节
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
		}))
		//构造消息

	})
}
