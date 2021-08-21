package main

import (
	"servercg/controllers"
	"servercg/utils"
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	//-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ Configuração inicial Iris/APP -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
	app := iris.Default()

	//-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ Diretório de requisição default -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
	app.HandleDir("/", "./public")

	//-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ Configuração da pasta com as Views -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	// -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_- Rotas -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
	app.Get("/", controllers.WeAreInTheAr())

	app.Post("/record/insert", controllers.Insert())

	app.Post("/record/update", controllers.Update())

	app.Get("/record", controllers.List())

	//-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ Configuração de porta e o ip/ Start app -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
	app.Run(iris.Addr(fmt.Sprint(":", utils.Godotenv("port_application"))))
}


