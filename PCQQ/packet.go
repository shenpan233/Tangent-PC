/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		组包头尾
* @Creat:   2021/11/27 0027 12:31
 */
package PCQQ

import "Tangent-PC/utils/GuBuffer"

func (this *TangentPC) packetLogin(cmd uint16, bin []byte) (uint16, []byte) {
	seq := this.udper.GetSeq()
	return seq, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(2)
		pack.SetUint16(this.sdk.CMainVer)
		pack.SetUint16(cmd)
		pack.SetUint16(seq)
		pack.SetUint32(uint32(this.info.LongUin))
		pack.SetUint8(03)
		pack.SetUint16(0)
		pack.SetUint32(0x00_01_01_01)
		pack.SetUint32(this.sdk.DwPubNo)
		pack.SetUint32(0)
		pack.SetBytes(bin)
		pack.SetUint8(3)
	})
}
