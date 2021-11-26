package GuBuffer

import (
	util "Tangerine/utils"
	"bytes"
)

func NewGuPacket() (p *GuPacket) {
	p = new(GuPacket)
	p.w = new(bytes.Buffer)
	return
}
func (p *GuPacket) Reset() {
	p.w = new(bytes.Buffer)
}
func (p *GuPacket) JmpHead() {
	p.tmp = p.w.Bytes()
	p.w.Reset()
}
func (p *GuPacket) SetHex(hex string) {
	p.SetBytes(util.HexToBin(hex))
}
func (p *GuPacket) SetBytes(bin []byte) {
	p.w.Write(bin)
}
func (p *GuPacket) SetToken(bin []byte) {
	p.SetInt16(int16(len(bin)))
	p.SetBytes(bin)
}
func (p *GuPacket) SetSToken(bin string) {
	p.SetInt16(int16(len(bin)))
	p.SetBytes([]byte(bin))
}

func (p *GuPacket) SetString(bin string) {
	p.w.WriteString(bin)
}
func (p *GuPacket) SetTlv(t *Tlv) {
	p.SetInt16(int16(t.Tag))
	p.SetInt16(int16(len(t.Value)))
	p.w.Write(t.Value)
}
func (p *GuPacket) SetInt8(i int8) {
	p.w.Write([]byte{(byte)(i >> 0)})
}
func (p *GuPacket) SetInt16(i int16) {
	p.w.Write([]byte{(byte)(i >> 8), (byte)(i >> 0)})
}
func (p *GuPacket) SetInt32(i int32) {
	p.w.Write([]byte{(byte)(i >> 24), (byte)(i >> 16), (byte)(i >> 8), (byte)(i >> 0)})
}
func (p *GuPacket) SetInt64(i int64) {
	p.w.Write([]byte{(byte)(i >> 56), (byte)(i >> 48), (byte)(i >> 40), (byte)(i >> 32), (byte)(i >> 24), (byte)(i >> 16), (byte)(i >> 8), (byte)(i >> 0)})
}
func (p *GuPacket) GetAll() (bin []byte) {
	p.w.Write(p.tmp)
	p.tmp = nil
	bin = p.w.Bytes()
	return
}

func (p *GuPacket) GetHex() string {
	return util.BinToHex(p.GetAll())
}

func (p *GuPacket) ToTlv(t int16) (bin []byte) {
	bin = p.GetAll()
	p.Reset()
	p.SetInt16(t)
	p.SetInt16(int16(len(bin)))
	p.SetBytes(bin)
	return p.GetAll()
}
