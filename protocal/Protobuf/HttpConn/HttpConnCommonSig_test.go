/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/4 12:38
  @Notice:  notice
*/

package HttpConn

import (
	"testing"
)

func TestPack_HttpConn_sig(t *testing.T) {
	//fmt.Println(util.BinToHex())
}

func TestPost(t *testing.T) {
	//r := GetReqBody(GetHttpConnSig(256783314, 55, model.Version{
	//	DwSSOVersion: 0x00_00_04_5C,
	//	DwPubNo:      0x00_00_6A_0A,
	//	ServiceId:    0x00_00_00_01,
	//	ClientVer:    0x00_00_16_BD,
	//	CMainVer:     0x3A_15,
	//	DwQdVersion:  0x04_05_00_09,
	//	DwClientType: 0x00_01_01_01,
	//	ClientMd5:    []byte{0xD0, 0x1D, 0x63, 0xA5, 0x85, 0x28, 0x01, 0x97, 0x59, 0x8C, 0xEC, 0xFF, 0x29, 0xC6, 0x31, 0xA3},
	//}, util.HexToBin("AA 2E DF E8 82 D7 20 49 AC 98 EB E6 09 EC 92 70 77 63 96 B8 A0 3F D2 16 BF C1 10 20 D8 5D 3F 86 A0 D3 A0 F5 D2 6F 30 0B 7C 95 44 22 6C 15 9B 8F AC E5 49 B1 0F 6E 07 96 9F DF 47 ED 12 13 16 07 61 5D 6B 45 17 DD C6 D6 45 7A DB F7 9E 7A 0E 51 80 B2 91 64 F4 1B 0B B4 57 9C 02 8C 7A 59 09 66 1F BC 3F 40 AC 0D 24 96")), "Zpm6ilscj7")

	//client := req.C().EnableKeepAlives().EnableDumpAll()
	//content := GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
	//	pack.SetUint8(0x28)
	//	pack.SetUint32(uint32(len(r)))
	//	pack.SetUint32(24)
	//	pack.SetBytes(r)
	//	pack.SetBytes(util.GetRandomBin(24))
	//	pack.SetUint8(0x29)
	//})
	//post, err := client.R().SetHeaders(map[string]string{
	//	"Accept":          "*/*",
	//	"User-Agent":      "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
	//	"Connection":      "Keep-Alive",
	//	"Cache-Control":   "no-cache",
	//	"Accept-Encoding": "gzip, deflate",
	//}).SetContentType("application/octet-stream").SetBody(content).Post("http://120.232.130.90/cgi-bin/httpconn")
	//fmt.Println(post, err)
}
