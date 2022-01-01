/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		协议选择器
* @Creat:   2021/11/26 0026 23:00
 */
package PCQQ

import (
	"Tangent-PC/model"
	util "Tangent-PC/utils"
)

func (this *TangentPC) U948() {
	this.sdk = &model.Version{
		DwSSOVersion: 0x00_00_04_5C,
		DwPubNo:      0x00_00_6A_0A,
		ServiceId:    0x00_00_00_01,
		ClientVer:    0x00_00_16_BD,
		CMainVer:     0x3A_15,
		ClientMd5:    util.HexToBin("58 D1 6F C1 80 EF 1F F1 0D FA 98 98 F8 DF 75 AD"),
	}
	this.sdk.CSubVer = this.sdk.CMainVer
}
