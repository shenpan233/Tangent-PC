# Author:Trial
# 输出打印工具
# email:shenpan233@vip.qq.com
# CreateDate: 2021/11/26 0026 20:04

class context:
    def __init__(self, Msg, tips=""
                                 ""):
        self.msg = Msg
        self.tips = tips
        return

    def print(self):
        msg = self.msg
        if self.tips != "":
            msg += "  //" + self.tips
        return msg


class outPrint:
    def __init__(self):
        self.tree = []
        return

    def Add(self):
        self.tree.append(context(1, "2", "3"))
