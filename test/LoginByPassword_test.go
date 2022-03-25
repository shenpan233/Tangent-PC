/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/5 16:48
  @Notice:  账号密码登录测试
*/

package test

import "testing"

func TestTgTLogin(t *testing.T) {
	if client.PingServer() {
		client.Login("2323232")
	}
}
