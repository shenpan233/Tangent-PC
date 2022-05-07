/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 21:02
  @Notice:  图片信息
*/

package FriendMsg

import (
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

type Img struct {
	Guid string
}

func (i *Img) Generate() []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetLitTlv(0x06, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetLitTlvS(0x02, "{"+i.Guid+"}.png")
			pack.SetLitTlv(0x04, util.GetRandomBin(4))
			pack.SetLitTlv(0x05, util.GetRandomBin(4))
			pack.SetLitTlv(0x06, []byte{0x00, 0x00, 0x00, 0x50}) //这个居然是固定的，应该什么乱七八糟的参数一样吧
			pack.SetLitTlv(0x07, []byte{0x43})                   //这个居然是固定的，应该什么乱七八糟的参数一样吧
			pack.SetLitTlv(0x08, nil)
			pack.SetLitTlv(0x09, []byte{0x01}) //这个居然是固定的，应该什么乱七八糟的参数一样吧
			pack.SetLitTlv(0x0A, util.Guid2Md5Bytes(i.Guid))
			pack.SetLitTlv(0x0B, nil)
			pack.SetLitTlv(0x14, []byte{0x00, 0x00, 0x00, 0x00})
			pack.SetLitTlv(0x15, Bytes.Uint32ToBytes(1080)) //TODO 发送图片的宽度
			pack.SetLitTlv(0x16, Bytes.Uint32ToBytes(1080)) //TODO 发送图片的高度
			pack.SetLitTlv(0x18, Bytes.Uint32ToBytes(100))  //TODO 发送图片的大小bit
			pack.SetLitTlv(0x1C, Bytes.Uint32ToBytes(1001))
			pack.SetLitTlv(0xFF, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
				pack.SetUint16(1536) //这个是固定的
				pack.SetBytes([]byte{0x20})
				pack.SetString("92kA1C9d30d2732920e878")
				//pack.SetBytes([]byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x20})
				//pack.SetBytes([]byte{0x35, 0x30})
				//pack.SetBytes([]byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20})
				//性能优化
				pack.SetBytes([]byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x35, 0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20})
				pack.SetString(i.Guid)
				pack.SetString("A")
			}))

		}))
	})
}

func (i *Img) DecodeString() string {
	return i.Guid
}
