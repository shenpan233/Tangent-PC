/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		协议选择器
* @Creat:   2021/11/26 0026 23:00
 */

package PCQQ

import (
	"github.com/shenpan233/Tangent-PC/model"
	util "github.com/shenpan233/Tangent-PC/utils"
)

func (this *TangentPC) U948() {
	this.sdk = &model.Version{
		DwSSOVersion: 0x00_00_04_5C,
		DwPubNo:      0x00_00_6A_0A,
		ServiceId:    0x00_00_00_01,
		ClientVer:    0x00_00_16_BD,
		CMainVer:     0x3A_15,
		DwQdVersion:  0x04_05_00_09,
		ClientMd5:    util.HexToBin("D01D63A585280197598CECFF29C631A3"),
	}
	this.sdk.CSubVer = this.sdk.CMainVer
}
