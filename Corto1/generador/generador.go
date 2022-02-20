package generador

import "strconv"

var tmp = 0

func NewTemp() string {
	tmp = tmp + 1
	return "t" + strconv.Itoa(tmp)
}
