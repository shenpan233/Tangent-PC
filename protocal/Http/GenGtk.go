package Http

import "strconv"

//GenGtk Gtk计算
func GenGtk(skey string) string {
	hash := 5381
	for i := 0; i < len(skey); i++ {
		hash += (hash << 5) + int(skey[i])
	}
	return strconv.Itoa(hash & 0x7fffffff)
}
