/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/4 12:38
  @Notice:  notice
*/

package HttpConn

import (
	"fmt"
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
	"testing"
)

func TestPack_HttpConn_sig(t *testing.T) {
	//fmt.Println(util.BinToHex())
}

func TestPost(t *testing.T) {
	r := GetReqBody(GetHttpConnSig(256783314, 55, model.Version{
		DwSSOVersion: 0x00_00_04_5C,
		DwPubNo:      0x00_00_6A_0A,
		ServiceId:    0x00_00_00_01,
		DwAppVer:     0x00_00_16_BD,
		CMainVer:     0x3A_15,
		DwQdVersion:  0x04_05_00_09,
		DwClientType: 0x00_01_01_01,
		ClientMd5:    []byte{0xD0, 0x1D, 0x63, 0xA5, 0x85, 0x28, 0x01, 0x97, 0x59, 0x8C, 0xEC, 0xFF, 0x29, 0xC6, 0x31, 0xA3},
	}, util.HexToBin("D1 BF BF 42 8E F5 28 18 D7 D9 EB 2E 07 BB 36 51 F0 10 23 6A AF 35 13 47 64 58 00 0C 3F 60 D2 ED 30 DD C7 C2 1E 0E 2E 32 8B 83 12 0D 8B 95 CE 8C C8 BF 8A 2E 50 85 FA 83 D0 94 6D F2 53 F6 B7 53 EF 63 21 F1 85 3E 11 79 21 EF 18 2F 41 C7 37 B1 F5 DF 92 4C 69 9A 35 3D A2 AF F1 2C F6 4B B7 7F BD 82 29 8F 1B CD 5C 6F")), "ZZECZEENcI")
	fmt.Println(util.BinToHex(r))
}
