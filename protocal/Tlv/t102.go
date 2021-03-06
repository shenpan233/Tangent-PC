/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Official
* @Creat:   2021/12/18 12:22
 */

package Tlv

import (
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/Bytes"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetTlv102Official(info *model.Information) []byte {
	OfficialSig := util.GetRandomBin(56)
	OfficialKey := util.GetRandomBin(16)
	pack := GuBuffer.NewGuPacket()
	pack.SetUint16(wSubVer)
	pack.SetBytes(CreateOfficial(info.PassWord, OfficialSig, OfficialKey))
	pack.SetToken(OfficialSig)
	pack.SetToken(GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetBytes(OfficialKey)
		pack.SetBytes(Bytes.GetCrc32(OfficialKey))
	}))
	return pack.ToTlv(0x102)
}
