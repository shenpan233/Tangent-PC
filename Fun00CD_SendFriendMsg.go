/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 19:30
  @Notice:  发送好友消息
*/

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func (this *TangentPC) pack00CD(FriendUin uint64, builder ...Msg.Builder) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_CD, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetUint32(uint32(FriendUin))
		pack.SetBytes([]byte{0, 0})
		pack.SetUint16(8)
		pack.SetBytes([]byte{0x00, 0x01})
		pack.SetUint16(4)
		pack.SetBytes([]byte{0x00, 0x00, 0x00, 0x00})
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetUint32(uint32(FriendUin))
		pack.SetBytes(Bytes.GetMd5Bytes(GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetUint32(uint32(FriendUin))
			pack.SetBytes(this.teaKey.SessionKey)
		})))
		pack.SetUint16(0x0B) //消息类型
		pack.SetBytes(util.GetRandomBin(2))
		pack.SetUint32(uint32(util.GetServerCurTime()))
		pack.SetBytes([]byte{0x00, 0x00})                                     //发送者头像???
		pack.SetBytes([]byte{0x00, 0x00, 0x00, 0x00})                         //字体样式属性
		pack.SetBytes([]byte{0x01, 0x00, 0x00, 0x00, 0x01})                   //一些分片信息 先不用管
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
		for _, f := range builder {
			pack.SetBytes(f.Generate())
		}
	}))
}

func (this *TangentPC) unpack00CD(bin []byte) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin), func(pack *GuBuffer.GuUnPacket) {

	})
}
