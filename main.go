package main

import (
	"reflect"

	"github.com/rew3/rew3-pkg/common/account"
)

func main() {
	model := account.MiniUser{}
	typ := reflect.TypeOf(model)
	println(typ.PkgPath())
	println(typ.Name())
}
