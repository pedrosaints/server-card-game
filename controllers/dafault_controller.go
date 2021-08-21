package controllers

import (
	"github.com/kataras/iris/v12"
)

func WeAreInTheAr() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.WriteString("<h1>TAMO ONLINE</h1>")
	}
}



