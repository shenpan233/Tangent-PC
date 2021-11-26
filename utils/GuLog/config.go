/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		全局初始化
* @Creat:   2021/10/29 0029 22:43
 */
package GuLog

//Config 初始化配置
//	ForceColor 是否强制彩色模式
//	formatTime 时间格式化,如果为空默认为"2006/01/02 15:04:05.0000"
func Config(ForceColor bool, formatTime string) {
	guLog.color = ForceColor
	if formatTime == "" {
		guLog.timeFormat = "2006/01/02 15:04:05.0000"
	} else {
		guLog.timeFormat = formatTime
	}
}
