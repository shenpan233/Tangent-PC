/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Main测试
* @Creat:   2021/11/26 0026 22:44
 */
package demo

import (
	client2 "Tangent-PC/PCQQ"
	"Tangent-PC/utils/GuLog"
	"testing"
)

func TestBuild1(t *testing.T) {
	GuLog.Config(true, "")
	client := client2.New("0")
	client.U948()

	client.Ping()

	select {}

}
