/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/3/19 14:39
  @Notice:  公用文件
*/

package Http

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39"

func PkgCommonCookies(uin string, skey, baseKey string) map[string]string {
	return map[string]string{
		"p_uin": "o0" + uin, "uin": "o0" + uin, "p_skey": baseKey, "skey": skey,
	}
}
