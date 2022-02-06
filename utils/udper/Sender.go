/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		发送
* @Creat:   2021/11/27 0027 13:10
 */
package udper

import "time"

func (this *Udper) Send(bin *[]byte) bool {
	if _, err := this.conn.Write(*bin); err != nil {
		return true
	} else {
		return false
	}
}

//	SsoSeq	封包序列
//	WaitTime 等待时间,单位秒
func (this *Udper) SendAndGet(SsoSeq uint16, WaitTime uint32, bin *[]byte) []byte {
	if _, exist := this.pull.Load(SsoSeq); exist {
		panic("相同的SsoSeq? 发生什么事了?")
	} else {
		puller := make(chan []byte) //申请channel
		this.Send(bin)
		this.pull.Store(SsoSeq, puller)
		select {
		case bin := <-puller:
			close(puller)
			this.pull.Delete(SsoSeq)
			return bin
		case <-time.After(time.Duration(WaitTime) * time.Second):
			/*超时删除*/
			this.pull.Delete(SsoSeq)
			return nil
		}

	}
}
