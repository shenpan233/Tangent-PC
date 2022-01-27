/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:     SSo的信息
* @Creat:   2021/12/4 0004 14:57
 */

package Tlv

import (
	"Tangent-PC/model"
	"Tangent-PC/utils/GuBuffer"
)

func GetTlv19SSOInfo(sdk *model.Version) []byte {
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetUint32(sdk.DwSSOVersion)
	pack.SetUint32(sdk.ServiceId)
	pack.SetUint32(sdk.ClientVer)
	pack.SetUint16(wUnknown1)
	return pack.ToTlv(0x00_19)
}
