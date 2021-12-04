/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Main测试
* @Creat:   2021/11/26 0026 22:44
 */
package demo

import (
	client2 "Tangent-PC/PCQQ"
	"Tangent-PC/model"
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuLog"
	"fmt"
	"testing"
)

func TestBuild1(t *testing.T) {
	GuLog.Config(true, "")
	client := client2.New("0", model.Computer{
		util.HexToBin("9B 13 EE DA 00 00 00 00 00 00 00 00 00 00 00 00"),
		util.HexToBin("3B AE C1 AD 3C 07 44 EE 29 BE C0 38 4E F2 4A 1A"),
	})
	client.U948()

	if client.PingServer() {
		GuLog.Notice("System", "[QQ=%s]Ping成功", "0")
		fmt.Println(client.FetchQRCode().QRCode)

	}

	select {}

}
