/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/3 22:20
  @Notice:  事件处理器
*/

package Tangent_PC

func (this *TangentPC) eventHandles() map[uint16]unpack {
	return map[uint16]unpack{
		0x00_17: this.unpack0017,
		0x00_CE: this.unpack00CE,
	}
}
