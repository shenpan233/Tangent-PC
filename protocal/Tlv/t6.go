/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		密码验证TGTGT
* @Creat:   2021/12/18 10:48
 */

package Tlv

import (
	"Tangent-PC/model"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv6TGTGT(QQInfo *model.Information, version *model.Version, TGTKey []byte) []byte {
	pack := GuBuffer.NewGuPacket()
	Password2 := util.ToMd5Bytes(QQInfo.PassWord)
	pack.SetBytes(util.Encrypt(Password2, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint32(uint32(util.GetRand32()))
		pack.SetUint16(2)
		pack.SetUint32(uint32(QQInfo.LongUin))
		pack.SetUint32(version.DwSSOVersion)
		pack.SetUint32(version.ClientVer)
		pack.SetUint32(uint32(version.CMainVer))
		pack.SetUint16(0)
		pack.SetUint8(1) //是否记住密码
		pack.SetBytes(QQInfo.PassWord)
		pack.SetUint32(QQInfo.PingTime)
		pack.SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		pack.SetUint32(uint32(util.IpToInt(QQInfo.WlanIp)))
		pack.SetBytes([]byte{0, 0, 0, 0, 0, 0})
		pack.SetTlv(&GuBuffer.Tlv{
			Tag:   0x00_00,
			Value: QQInfo.ComputerIdEx,
		})
		pack.SetBytes(TGTKey)
	})))
	return pack.ToTlv(0x6)
}
