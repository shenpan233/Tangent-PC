/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/19 14:39
  @Notice:  公用文件
*/

package Http

import (
	"net/http"
)

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39"

func PkgCommonCookies(uin string, skey, baseKey string) (cookies []*http.Cookie) {
	cookies = append(cookies, &http.Cookie{
		Name:  "p_uin",
		Value: "o0" + uin,
	}, &http.Cookie{
		Name:  "uin",
		Value: "o0" + uin,
	}, &http.Cookie{
		Name:  "p_skey",
		Value: baseKey,
	}, &http.Cookie{
		Name:  "skey",
		Value: skey,
	})
	return
}
