/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 19:25
  @Notice:  发送消息
*/

package Send

import (
	"github.com/golang/protobuf/proto"
	"github.com/shenpan233/Tangent-PC/model"
	Model "github.com/shenpan233/Tangent-PC/protocal/Msg/Group"
	"github.com/shenpan233/Tangent-PC/protocal/Protobuf/im/cmd0x0002"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"reflect"
)

//GroupMsg 发送群消息 废弃已重构
//  @deprecated
//	Msg 消息内容
func GroupMsg(GroupUin uint64, Msg string, atName GetGroupMemberCardFromCache) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(Model.Send) //事件类型
		pack.SetUint32(uint32(GroupUin))
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
			ret := BuildMsgStructure(Msg, GroupUin, atName)
			for _, subCall := range ret {
				bin := reflect.ValueOf(subCall).MethodByName("Marshal").Call(nil)[0].Bytes()
				pack.SetBytes(bin)
			}

		}))

	})
}

//Recall 发送消息回调
func Recall(bin []byte) (isSuc bool, Recall cmd0x0002.SendGroupMsg) {
	pack := GuBuffer.NewGuUnPacket(bin)
	isSuc = pack.GetUint8() == model.LogicSuc
	if !isSuc {
		pack.Skip(4) //MsgRandom好像没用
		_ = proto.Unmarshal(pack.GetToken(), &Recall)
	}
	return
}
