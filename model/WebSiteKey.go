/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/3 12:48
  @Notice:  网页操作相关Key
*/

package model

const (
	WebQun = "qun.qq.com"
)

var (
	WebSite = []string{"t.qq.com", WebQun, "qzone.qq.com", "qzone.qq.com", "ke.qq.com"}
)

type WebKey struct {
	Skey, PSkey string
	//WebSiteKeys
	//	BufQRKey=website
	WebSiteKeys map[string]string
}
