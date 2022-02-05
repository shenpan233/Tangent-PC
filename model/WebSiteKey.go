/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/3 12:48
  @Notice:  网页操作相关Key
*/

package model

var (
	WebSite = []string{"t.qq.com", "qun.qq.com", "qzone.qq.com", "qzone.qq.com", "ke.qq.com"}
)

type WebKey struct {
	Common CommonWebKey
	//WebSiteKeys
	//	BufQRKey=website
	WebSiteKeys map[string]CommonWebKey
}

type CommonWebKey struct {
	Skey  string
	PSkey string
}
