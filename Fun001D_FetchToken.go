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
	subCmd0x001DHttpConn = 0x26
	subCmd0x001DWebKey   = 0x33 //网站的key
)

func (this *TangentPC) pack001D(t uint8) (SsoSeq uint16, buffer []byte) {
	return this.packetCommonEnc(0x00_1D, GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(t)
		switch t {
		case subCmd0x001DWebKey:
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
		case subCmd0x001DWebKey:
			key := model.WebKey{
				WebSiteKeys: make(map[string]string),
			}
			key.Skey = string(pack.GetToken())
			key.PSkey = string(pack.GetToken())
			num := int(pack.GetUint16()) - 1
			for i := 0; i < num; i++ {
				website := string(pack.GetToken())
				key.WebSiteKeys[website] = string(pack.GetToken())
			}
			ret = &key
			break
		case subCmd0x001DHttpConn:
			resp := make(map[int8][]byte)
			//61 ** ** ** 6A 68 ** ** 79 42 50 44 44 ** ** **
			resp[0] = pack.GetBin(16) //HttpConnTeaKey
			//FF FF FF FA
			//08 00 00 00 C8 00 40 40
			pack.Skip(12)
			//68
			//D1 BF BF 42 8E F5 28 18 D7 ** ** ** 07 BB 36 51 F0 10 23 6A AF ** ** ** 64 58 00 0C 3F 60 D2 ED 30 DD C7 C2 1E 0E 2E ** ** ** ** ** ** ** ** ** ** ** **  2E 50 85 FA 83 D0 94 6D F2 53 F6 B7 53 EF 63 ** ** ** 3E 11 79 21 EF 18 2F 41 C7 37 B1 F5 DF 92 4C ** ** ** 3D A2 AF F1 2C F6 4B B7 7F BD 82 29 8F 1B CD 5C 6F
			resp[1] = pack.GetBin(int(pack.GetUint8())) //BufSigHttpConnToken
			//13 80 02 04
			ret = resp
			break
		}
	})
	return
}
