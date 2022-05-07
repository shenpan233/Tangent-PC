/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/17 11:35
  @Notice:  notice
*/

package FriendMsg

import (
	"github.com/shenpan233/Tangent-PC/protocal/Msg"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

type Pic struct {
	Guid string `json:"guid"`
}

func (this *Pic) Generate() []byte {
	/*
		   	02 //子消息类型
		      00 2A //图片拖到桌面后显示的图片名字
		      7B 35 34 34 31 44 38 34 32 2D 35 46 35 42 2D 37
		      45 34 36 2D 43 33 30 39 2D 42 33 33 30 31 44 35
		      46 32 33 45 31 7D 2E 6A 70 67		[{5441D842-5F5B-7E46-C309-B3301D5F23E1}.jpg]

		      04	//Random
		      00 04
		      9D 30 D2 73


		      05	//Random
		      00 04
		      29 20 E8 78

		      06
		      00 04
		      00 00 00 50

		      07
		      00 01
		      43

		      08
		      00 00

		      09
		      00 01
		      01

		      0A //图片MD5 吧guid拆开就行
		      00 10
		      54 41 D8 42 5F 5B 7E 46 C3 09 B3 30 1D 5F 23 E1

		      0B
		      00 00

		      14
		      00 04
		      00 00 00 00

				  15
		      00 04
		      00 00 00 7A //宽度


		      16
		      00 04
		      00 00 00 38 //高度

		      18
		      00 04
		      00 00 03 B8 //图片大小

		      1C
		      00 04
		      00 00 03 E9 [不清楚:1001]

		      FF
		      00 5C
		      15 36 [不清楚 5430]
		      20	//分隔符
		      39 32 6B 41 31 43 39 64 33 30 64 32 37 33 32 39
		      32 30 65 38 37 38 [92kA1C9d30d2732920e878]
		      20 20 20 20 20 20
		      35
		      30
		      20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20
		      7B 35 34 34 31 44 38 34 32 2D 35 46 35 42 2D 37
		      45 34 36 2D 43 33 30 39 2D 42 33 33 30 31 44 35
		      46 32 33 45 31 7D 2E 6A 70 67 [{5441D842-5F5B-7E46-C309-B3301D5F23E1}.jpg]
		      41 [文件属性:A]
	*/
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetLitTlv(Msg.TypePic, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetLitTlv(0x02, []byte(this.Guid))
			pack.SetLitTlv(0x04, util.GetRandomBin(4))
			pack.SetLitTlv(0x05, util.GetRandomBin(4))
			pack.SetLitTlv(0x06, []byte{0x00, 0x00, 0x00, 0x50}) //这个居然是固定的，应该什么乱七八糟的参数一样吧
			pack.SetLitTlv(0x07, []byte{0x43})                   //这个居然是固定的，应该什么乱七八糟的参数一样吧
			pack.SetLitTlv(0x08, nil)
			pack.SetLitTlv(0x09, []byte{0x01}) //这个居然是固定的，应该什么乱七八糟的参数一样吧
			pack.SetLitTlv(0x0A, util.Guid2Md5Bytes(this.Guid))
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
				pack.SetString(this.Guid)
				pack.SetString("A")
			}))

		}))
	})
}

func (p *Pic) Decode() string {
	return Msg.FormatPic + p.Guid + Msg.FormatEnd
}
