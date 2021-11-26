# Author:Trial
# email:shenpan233@vip.qq.com
# 分析小工具
# CreateDate: 2021/11/26 0026 18:56
# 我的Python是真的差
import out
from BufReader import BufReader

if __name__ == '__main__':
    sendto = input()
    try:
        buf = BufReader(sendto)
        head, headHex = buf.GetUintFromLen(1, True)
        if head == 2:
            buf = BufReader(buf.GetAllHex())
            sendto = buf.GetHex(buf.Len() - 1)
            if buf.GetUintFromLen(1) == 3:
                buf = BufReader(sendto)
                print(out.context(headHex, "包头").print())
                i, Hex = buf.GetUintFromLen(2, True)
                print(out.context(Hex, "Version[%d]" % i).print())

                cmd, Hex = buf.GetUintFromLen(2, True)
                print(out.context(Hex, "包命令[%d]" % cmd).print())

                i, Hex = buf.GetUintFromLen(2, True)
                print(out.context(Hex, "Seq[%d]" % i).print())

                i, Hex = buf.GetUintFromLen(4, True)
                print(out.context(Hex).print())

                i, Hex = buf.GetUintFromLen(3, True)
                print(out.context(Hex).print())

                i, Hex = buf.GetUintFromLen(4, True)
                print(out.context(Hex).print())

                i, Hex = buf.GetUintFromLen(4, True)
                print(out.context(Hex).print())

                i, Hex = buf.GetUintFromLen(4, True)
                print(out.context(Hex, "QQUin[%d]" % i).print())

                if cmd == 0x0819:
                    buf.GetTlv()

                print(out.context(buf.GetAllHex()).print())
                print(out.context("03", "包尾").print())
    except BaseException:
        print("待分析数据异常")
        pass
