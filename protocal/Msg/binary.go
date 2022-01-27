/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/27 19:59
  @Notice:  消息实例构造器
*/

package Msg

import (
	"Tangent-PC/utils/GuBuffer"
	"fmt"
)

type (
	//接口
	message interface {
		ToString() string
		Marshal() []byte
	}
	Common struct {
		IsAt  bool
		AtUin uint32
		Msg   string
		message
	}
)

func (this *Common) ToString() string {
	if this.IsAt {
		if this.AtUin == 0 {
			return FormatAt + "All" + FormatEnd
		} else {
			return fmt.Sprintf(FormatAt+"%d"+FormatEnd, this.AtUin)
		}
	}
	return this.Msg
}

func (this *Common) Marshal() []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetLitTlv(TypeText, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			if this.IsAt {
				this.Msg = "@小可爱" //TODO 这里At获取名称没搞定
			}
			pack.SetLitTlv(CommonMsg, []byte(this.Msg))
			if this.IsAt {
				pack.SetLitTlv(CommonAt, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
					pack.SetBytes([]byte{0x00, 0x01})
					pack.SetUint32(4) //TODO 这个是上面this.Msg的字符长度 就是肉眼能看到的长度
					pack.SetUint32(this.AtUin)
					pack.SetBytes([]byte{0x00, 0x00})
				}))
			}
		}))
	})
}

type (
	Pic struct {
		Guid string
		message
	}
)

func (this *Pic) ToString() string {
	return FormatPic + this.Guid + FormatEnd
}

func (this *Pic) Marshal() []byte {
	/*
	   	02 //子消息类型
	      00 2A //图片拖到桌面后显示的图片名字
	      7B 35 34 34 31 44 38 34 32 2D 35 46 35 42 2D 37
	      45 34 36 2D 43 33 30 39 2D 42 33 33 30 31 44 35
	      46 32 33 45 31 7D 2E 6A 70 67
	      [
	      {5441D842-5F5B-7E46-C309-B3301D5F23E1}.jpg
	      ] //图片的MD5转UUID {uuid}.png
	      04
	      00 04
	      9D 30 D2 73


	      05
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
	      20
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
		pack.SetLitTlv(TypePic, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
			pack.SetLitTlv(0x02, []byte(this.Guid))
			pack.SetLitTlv(0x04, nil)
		}))
	})
}
