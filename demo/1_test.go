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
	"os"
	"testing"
	"time"
)

func TestBuild1(t *testing.T) {

	GuLog.Config(true, "")
	client := client2.New("0", model.Computer{
		ComputerId:   util.HexToBin("9B 13 EE DA 00 00 00 00 00 00 00 00 00 00 00 00"),
		ComputerIdEx: util.HexToBin("3B AE C1 AD 3C 07 44 EE 29 BE C0 38 4E F2 4A 1A"),
		DeviceID:     util.GetRandomBin(32),
		ComputerName: "Beijing University",
		MacGuid:      util.GetRandomBin(16),
	})
	if client == nil {
		return
	}
	client.U948()

	if client.PingServer() {
		GuLog.Notice("System", "[QQ=%s]Ping成功", "0")
		resp := client.FetchQRCode()
		os.WriteFile("./QRCode.png", resp.QRCode, os.FileMode(0777))
		go func() {
			for client.CheckQRCode(resp) != client2.QROk {
				time.Sleep(3 * time.Second)
			}
			client.QRLogin()
		}()
	}

	select {}

}
