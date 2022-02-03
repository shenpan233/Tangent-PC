/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/1 17:46
  @Notice:	刷新token
*/

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

const (
	webKey = 0x33 //网站的key
)

func (this *TangentPC) pack001D(t uint8) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_1D, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(t)
		switch t {
		case webKey:
			pack.SetUint16(uint16(len(model.WebSite)))
			for _, website := range model.WebSite {
				pack.SetSToken(website)
			}
			break
		}
	}))
}

func (this *TangentPC) unpack001D(bin []byte) (ret interface{}) {
	GuBuffer.NewGuUnPacketFun(util.Decrypt(this.teaKey.SessionKey, bin[3:]), func(pack *GuBuffer.GuUnPacket) {
		subCmd := pack.GetUint8()
		pack.Skip(1) //业务代码，一般都会成功懒得搞
		switch subCmd {
		case webKey:
			key := model.WebKey{}
			key.Common.Skey = string(pack.GetToken())
			key.Common.PSkey = string(pack.GetToken())
			num := int(pack.GetUint16())
			for i := 0; i < num; i++ {
				website := string(pack.GetToken())
				key.WebSiteKeys[website] = model.CommonWebKey{
					Skey:  string(pack.GetToken()),
					PSkey: string(pack.GetToken()),
				}
			}
			ret = &key
			break
		}
	})
	return
}
