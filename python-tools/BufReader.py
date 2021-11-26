# Author:Trial
# email:shenpan233@vip.qq.com
# CreateDate: 2021/11/26 0026 19:08


class BufReader:
    def __init__(self, hexData=""):
        if hexData != "":
            self.bin = bytes.fromhex(hexData.replace(" ", ""))
            self.offsite = 0
            self.size = len(self.bin)
        return

    def GetUint16(self, PrintHex=False):
        s, l = self.increase(2)
        data = self.bin[s:l]
        if PrintHex:
            return int.from_bytes(data[:2], "big"), self.spSpace(data.hex().upper())
        else:
            return int.from_bytes(data[:2], "big")

    def GetUintFromLen(self, length, PrintHex=False):
        s, l = self.increase(length)
        data = self.bin[s:l]
        if PrintHex:
            return int.from_bytes(data[:length], "big"), self.spSpace(data.hex().upper())
        else:
            return int.from_bytes(data[:length], "big")

    def GetBin(self, length, PrintHex=False):
        s, l = self.increase(length)
        data = self.bin[s:l]
        if PrintHex:
            return data[:length], self.spSpace(data.hex().upper())
        else:
            return data[:length]

    def GetHex(self, length):
        s, l = self.increase(length)
        data = self.bin[s:l]
        return self.spSpace(data.hex().upper())

    def GetTlv(self):
        _, TagHead = self.GetUintFromLen(2, True)
        length, lenHex = self.GetUintFromLen(2, True)
        val = self.GetHex(length)
        print(TagHead, lenHex, val)

    def Len(self):
        return len(self.bin)

    def GetAllHex(self):
        return self.bin[self.offsite:].hex().upper()

    # 偏移增长
    def increase(self, length):
        last = self.offsite
        offside = self.offsite + length
        if offside <= self.size:
            self.offsite = offside
            return last, self.offsite

    def spSpace(self, data=""):
        n = len(data)
        t = ""
        for i in range(0, n, 2):
            t = t + data[i:i + 2] + " "
        return t
