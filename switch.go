/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		协议选择器
* @Creat:   2021/11/26 0026 23:00
 */

package Tangent_PC

import (
	"github.com/shenpan233/Tangent-PC/model"
)

func (this *TangentPC) U948() *TangentPC {
	this.sdk = &model.Version{
		SAppName:     "QQ",
		CMainVer:     5821, //0x3A_15
		DwSSOVersion: 0x00_00_04_5C,
		DwPubNo:      0x00_00_6A_0A,
		ServiceId:    0x00_00_00_01,
		DwAppVer:     0x00_00_16_BD,
		DwQdVersion:  0x04_05_00_09,
		DwClientType: 0x00_01_01_01,
		DwSSOConfig:  0x0F_57_A2_48,
		ClientMd5:    []byte{0xD0, 0x1D, 0x63, 0xA5, 0x85, 0x28, 0x01, 0x97, 0x59, 0x8C, 0xEC, 0xFF, 0x29, 0xC6, 0x31, 0xA3},
	}
	this.sdk.CSubVer = this.sdk.CMainVer
	return this
}
