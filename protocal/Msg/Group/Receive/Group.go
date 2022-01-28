/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/2 20:53
  @Notice:  群消息接收处理
*/

package Receive

import (
	"bytes"
	"fmt"
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	"github.com/shenpan233/Tangent-PC/protocal/Msg/Group"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
	"github.com/shenpan233/Tangent-PC/utils/GuLog"
	"strings"
)

func GroupMsg(data []byte) (Msg model.GroupMsg) {
	GuBuffer.NewGuUnPacketFun(data, func(pack *GuBuffer.GuUnPacket) {
		{
			GuBuffer.NewGuUnPacketFun(pack.GetBin(int(pack.GetUint32())), func(pack *GuBuffer.GuUnPacket) {
				//Part1
			})
		}
		//第二部分
		{
			Msg.GroupUin = uint64(pack.GetUint32()) //群号
			pack.Skip(1)
			{
				Msg.SenderUin = uint64(pack.GetUint32())       //发消息的人
				Msg.MsgSeq = pack.GetUint32()                  //消息Seq
				Msg.MsgTime.Receive = uint64(pack.GetUint32()) //消息的时间
			}
			pack.Skip(8)
			//piceceNum := pack.GetUint8()
			//piceceId := pack.GetUint8()
			//piceceKind := pack.GetInt16()
			pack.Skip(4)
			pack.Skip(12)
			Msg.MsgTime.Send = uint64(pack.GetUint32()) //发送的时间
			Msg.MsgID = pack.GetUint32()
			pack.Skip(1)
			{
				Msg.Red = pack.GetUint8()
				Msg.Blue = pack.GetUint8()
				Msg.Green = pack.GetUint8()
				Msg.Size = pack.GetUint8()
				Msg.Style = pack.GetUint8()
				Msg.Encoding = uint16(pack.GetInt16())
				Msg.Font.FontName = pack.GetStr(int32(pack.GetInt16()))
				pack.Skip(2)
			}
		}
		Msg.Msg = groupMsgUnpack(pack)
	})
	fmt.Println(Msg)
	return
}

//groupMsgUnpack 解包群消息并结构化
func groupMsgUnpack(pack *GuBuffer.GuUnPacket) string {
	msgBuilder := bytes.NewBuffer(nil)
	for pack.GetLen() > 0 {
		MsgType := pack.GetUint8()
		GuBuffer.NewGuUnPacketFun(pack.GetToken(), func(pack *GuBuffer.GuUnPacket) {
			switch MsgType {
			case Msg.TypeText:
				switch pack.GetUint8() {
				case Msg.CommonMsg:
					//检查是否有At的消息
					common := Group.Common{
						Msg: string(pack.GetToken()),
					}
					GuBuffer.NewGuUnPacketFun(pack.GetAll(), func(pack *GuBuffer.GuUnPacket) {
						pack.Skip(1)
						pack = GuBuffer.NewGuUnPacket(pack.GetBin(int(pack.GetInt16())))
						if pack.GetLen() != 0 {
							//有其他的内容
							pack.Skip(7)
							common.IsAt = true
							common.AtUin = pack.GetUint32()
						}
					})
					msgBuilder.WriteString(common.ToString())
					break
				}
				break
			case Msg.TypeFace:
				pack.Skip(1)
				pack.Skip(2)
				msgBuilder.WriteString(strings.TrimSpace(pack.GetAllHex()))
				break
			case Msg.TypePic:
				tmp := pack.GetAll()
				pack = GuBuffer.NewGuUnPacket(tmp)
				pack.Skip(1)
				msgBuilder.WriteString(buildPic(string(pack.GetToken())))
				break
			default:
				GuLog.Warm("Group 解析", "Type=0x%d\nData=%X", MsgType, pack.GetAll())
			}
		})
	}
	return msgBuilder.String()
}
