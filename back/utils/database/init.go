package database

import (
	"reflect"
)

var currentPkg = ""

func init() {
	currentPkg = reflect.TypeOf(order).PkgPath()
}
