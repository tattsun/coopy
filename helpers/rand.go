package helpers

import (
	"github.com/dchest/uniuri"
)

func RandomStr(len int) string {
	return uniuri.NewLen(len)
}
