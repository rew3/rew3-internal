package main

import (
	"reflect"
	"github.com/rew3/rew3-internal/app/account"
)

func main() {
	model := account.MiniUser{}
	typ := reflect.TypeOf(model)
	println(typ.PkgPath())
	println(typ.Name())
}
